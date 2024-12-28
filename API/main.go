package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" 
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//تعریف مدل کارمند
type Employee struct {
	ID       int     `json:"id"`        // شناسه کارمند
	Name     string  `json:"name"`      // نام کارمند
	Email    string  `json:"email"`     // ایمیل کارمند
	Phone    string  `json:"phone"`     // شماره تلفن کارمند
	Position string  `json:"position"`  // سمت کارمند
	Salary   float64 `json:"salary"`    // حقوق کارمند
	Age      int     `json:"age"`       // سن کارمند
	TeamID   *int    `json:"team_id"`   // شناسه تیم (اختیاری)
	TeamName *string `json:"team_name"` // نام تیم (اختیاری)
}

//تعریف مدل تیم
type Team struct {
	ID          int     `json:"id"`           // شناسه تیم
	Name        string  `json:"name"`         // نام تیم
	Description string  `json:"description"`  // توضیحات تیم
	LeaderID    *int    `json:"leader_id"`    // شناسه رهبر تیم (اختیاری)
	LeaderName  *string `json:"leader_name"`  // نام رهبر تیم (اختیاری)
	MemberCount int     `json:"member_count"` // تعداد اعضای تیم
}

//متغیر برای ذخیره‌سازی اتصال به پایگاه داده
var db *sql.DB

//تابع اتصال به پایگاه داده
func Connect() error {
	var err error
	// اتصال به پایگاه داده MySQL. اطلاعات اتصال شامل کاربر، رمز عبور، میزبان و نام پایگاه داده می‌باشد.
	db, err = sql.Open("mysql", "admin:12345678@tcp(localhost:3306)/team_management_service")
	if err != nil {
		return err 
	}

	// بررسی اتصال به پایگاه داده
	err = db.Ping()
	if err != nil {
		return err 
	}

	log.Println("اتصال به پایگاه داده با موفقیت انجام شد")
	return nil
}

func main() {
	//اتصال به پایگاه داده هنگام شروع برنامه
	if err := Connect(); err != nil {
		log.Fatal("خطا در اتصال به پایگاه داده:", err)
	}

	//ایجاد برنامه با استفاده از فایبر
	app := fiber.New()

	//فعال کردن CORS برای تمام مسیرها
	app.Use(cors.New())

	//تعریف مسیرهای مدیریت کارمندان
	app.Get("/employees", getEmployees)          // دریافت تمام کارمندان
	app.Post("/employees", addEmployee)          // اضافه کردن کارمند جدید
	app.Put("/employees/:id", updateEmployee)    // به‌روزرسانی اطلاعات کارمند
	app.Delete("/employees/:id", deleteEmployee) // حذف کارمند

	//تعریف مسیرهای مدیریت تیم‌ها
	app.Get("/teams", getTeams)          // دریافت تمام تیم‌ها
	app.Post("/teams", addTeam)          // اضافه کردن تیم جدید
	app.Put("/teams/:id", updateTeam)    // به‌روزرسانی اطلاعات تیم
	app.Delete("/teams/:id", deleteTeam) // حذف تیم

	//تنظیم مدیریت خطاهای 500 برای برنامه فایبر
	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			log.Printf("خطای سرور داخلی: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "خطای سرور داخلی. لطفاً دوباره تلاش کنید.",
			})
		}
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}

//تابع برای دریافت تمام کارمندان
func getEmployees(c *fiber.Ctx) error {
	//کوئری برای دریافت اطلاعات کارمندان و تیم مرتبط
	query := `
	SELECT e.id, e.name, e.email, e.phone, e.position, e.salary, e.age, e.team_id, t.name AS team_name
	FROM employees e
	LEFT JOIN teams t ON e.team_id = t.id`

	//اجرای کوئری
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString("خطا در دریافت اطلاعات کارمندان")
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Phone, &emp.Position, &emp.Salary, &emp.Age, &emp.TeamID, &emp.TeamName)
		if err != nil {
			return c.Status(500).SendString("خطا در پردازش اطلاعات کارمندان")
		}
		employees = append(employees, emp)
	}

	return c.JSON(employees) //بازگشت اطلاعات کارمندان به صورت JSON
}

//تابع برای اضافه کردن کارمند جدید
func addEmployee(c *fiber.Ctx) error {
	emp := new(Employee)
	if err := c.BodyParser(emp); err != nil {
		return c.Status(400).SendString("داده‌های ورودی نامعتبر")
	}

	//کوئری برای درج اطلاعات کارمند جدید
	query := `
	INSERT INTO employees (name, email, phone, position, salary, age, team_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, emp.Name, emp.Email, emp.Phone, emp.Position, emp.Salary, emp.Age, emp.TeamID)
	if err != nil {
		return c.Status(500).SendString("خطا در اضافه کردن کارمند")
	}

	//دریافت شناسه کارمند جدید
	id, _ := result.LastInsertId()
	emp.ID = int(id)
	return c.Status(201).JSON(emp)
}

//تابع برای به‌روزرسانی اطلاعات کارمند
func updateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	emp := new(Employee)
	if err := c.BodyParser(emp); err != nil {
		return c.Status(400).SendString("داده‌های ورودی نامعتبر")
	}

	//کوئری برای به‌روزرسانی اطلاعات کارمند
	query := `
	UPDATE employees 
	SET name = ?, email = ?, phone = ?, position = ?, salary = ?, age = ?, team_id = ?
	WHERE id = ?`
	_, err := db.Exec(query, emp.Name, emp.Email, emp.Phone, emp.Position, emp.Salary, emp.Age, emp.TeamID, id)
	if err != nil {
		return c.Status(500).SendString("خطا در به‌روزرسانی اطلاعات کارمند")
	}

	return c.Status(200).JSON(emp)
}

//تابع برای حذف کارمند
func deleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id") //دریافت شناسه کارمند از مسیر

	//کوئری برای حذف کارمند
	_, err := db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString("خطا در حذف کارمند")
	}

	return c.SendString("کارمند حذف شد")
}

//تابع برای دریافت تمام تیم‌ها
func getTeams(c *fiber.Ctx) error {
	query := `SELECT t.id, t.name, t.description, t.leader_id, e.name AS leader_name, 
              (SELECT COUNT(*) FROM employees WHERE team_id = t.id) AS member_count
              FROM teams t
              LEFT JOIN employees e ON t.leader_id = e.id`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("خطا در اجرای کوئری: %v", err)
		return c.Status(500).SendString("خطا در دریافت اطلاعات تیم‌ها")
	}
	defer rows.Close()

	var teams []Team
	for rows.Next() {
		var team Team
		err := rows.Scan(&team.ID, &team.Name, &team.Description, &team.LeaderID, &team.LeaderName, &team.MemberCount)
		if err != nil {
			log.Printf("خطا در پردازش ردیف: %v", err)
			return c.Status(500).SendString("خطا در پردازش اطلاعات تیم‌ها")
		}
		teams = append(teams, team)
	}

	return c.JSON(teams)
}

//تابع برای اضافه کردن تیم جدید
func addTeam(c *fiber.Ctx) error {
	team := new(Team)
	if err := c.BodyParser(team); err != nil {
		return c.Status(400).SendString("داده‌های ورودی نامعتبر")
	}

	//کوئری برای درج تیم جدید
	query := `
	INSERT INTO teams (name, description, leader_id) 
	VALUES (?, ?, ?)`
	result, err := db.Exec(query, team.Name, team.Description, team.LeaderID)
	if err != nil {
		return c.Status(500).SendString("خطا در اضافه کردن تیم")
	}

	//دریافت شناسه تیم جدید
	id, _ := result.LastInsertId()
	team.ID = int(id)
	return c.Status(201).JSON(team)
}

//تابع برای به‌روزرسانی اطلاعات تیم
func updateTeam(c *fiber.Ctx) error {
	id := c.Params("id") 
	team := new(Team)
	if err := c.BodyParser(team); err != nil {
		return c.Status(400).SendString("داده‌های ورودی نامعتبر")
	}

	//کوئری برای به‌روزرسانی اطلاعات تیم
	query := `
	UPDATE teams 
	SET name = ?, description = ?, leader_id = ?
	WHERE id = ?`
	_, err := db.Exec(query, team.Name, team.Description, team.LeaderID, id)
	if err != nil {
		return c.Status(500).SendString("خطا در به‌روزرسانی تیم")
	}

	return c.Status(200).JSON(team)
}

//تابع برای حذف تیم
func deleteTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.Exec("DELETE FROM teams WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString("خطا در حذف تیم")
	}

	return c.SendString("تیم حذف شد")
}
