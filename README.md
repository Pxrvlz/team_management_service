## About The Project

![Product Screenshot](https://cdn.shopaccino.com/igmguru/articles/Learn-Golang.png?)

<div dir="rtl">

## معرفی
این پروژه یک سرویس مدیریت تیم‌ها است که با استفاده از زبان Go توسعه داده شده است. این سرویس شامل قابلیت‌های CRUD برای مدیریت اطلاعات کارمندان و تیم‌ها است. همچنین، یک رابط کاربری با استفاده از Vue.js برای تعامل با این سرویس در دسترس است.

## نیازمندی‌ها
برای نصب و استفاده از این برنامه، نیاز به نرم‌افزارهای زیر دارید:
- **Go 1.23 یا نسخه‌های بالاتر**
- **Node.js** و **npm** (برای اجرای رابط کاربری Vue.js)
- **MySQL**  به عنوان پایگاه داده

## مراحل نصب

### 1. نصب و راه‌اندازی دیتابیس

ابتدا باید پایگاه داده را راه‌اندازی کنید. از دستورات زیر برای ایجاد دیتابیس و جداول استفاده کنید:

```sql
-- ایجاد دیتابیس جدید برای مدیریت تیم‌ها
CREATE DATABASE IF NOT EXISTS team_management_service;

-- استفاده از دیتابیس ایجاد شده
USE team_management_service;

-- ایجاد جدول 'employees' برای ذخیره اطلاعات کارمندان
CREATE TABLE IF NOT EXISTS employees (
    id INT AUTO_INCREMENT PRIMARY KEY,           -- شناسه کارمند
    name VARCHAR(100) NOT NULL,                  -- نام کارمند
    email VARCHAR(100) UNIQUE NOT NULL,          -- ایمیل کارمند
    phone VARCHAR(20),                           -- شماره تماس کارمند
    position VARCHAR(50),                        -- سمت کارمند
    salary FLOAT NOT NULL,                       -- حقوق کارمند
    age INT NOT NULL,                            -- سن کارمند
    team_id INT                                   -- شناسه تیم (اختیاری)
);

-- ایجاد جدول 'teams' برای ذخیره اطلاعات تیم‌ها
CREATE TABLE IF NOT EXISTS teams (
    id INT AUTO_INCREMENT PRIMARY KEY,           -- شناسه تیم
    name VARCHAR(100) NOT NULL,                  -- نام تیم
    description TEXT,                            -- توضیحات تیم
    leader_id INT,                               -- شناسه رهبر تیم (کلید خارجی به جدول 'employees')
    FOREIGN KEY (leader_id) REFERENCES employees(id) ON DELETE SET NULL -- کلید خارجی
);
```

### 2. نصب سرویس Go

برای اجرای سرویس Go، ابتدا کد پروژه را کلون کنید:

```bash
git clone https://github.com/Pxrvlz/team_management_service.git
cd team_management_service
```

سپس dependenciesهای Go را نصب کنید:

```bash
go mod tidy
```

برای اجرای سرویس، از دستور زیر استفاده کنید:

```bash
go run main.go
```

این سرویس به طور پیش‌فرض بر روی پورت 3000 در دسترس خواهد بود.

### 3. نصب و راه‌اندازی رابط کاربری Vue.js

برای اجرای رابط کاربری Vue.js، ابتدا به دایرکتوری `ui` بروید:

```bash
cd ui
```

سپس، dependenciesهای Vue.js را نصب کنید:

```bash
npm install
```

برای شروع پروژه Vue.js:

```bash
npm run dev
```

### 4. پیکربندی اتصال به دیتابیس در کد Go

در فایل `main.go`، بخش اتصال به پایگاه داده با استفاده از MySQL تنظیم شده است. اطمینان حاصل کنید که اطلاعات اتصال به پایگاه داده شما صحیح است:

```go
db, err = sql.Open("mysql", "admin:12345678@tcp(localhost:3306)/team_management_service")
```

در صورتی که از اطلاعات دیگری برای اتصال به دیتابیس استفاده می‌کنید، این بخش را تغییر دهید.

### 5. استفاده از API

#### دریافت تمامی کارمندان (GET)
برای دریافت اطلاعات تمام کارمندان، یک درخواست GET به `/employees` ارسال کنید.

مثال:
```bash
curl http://localhost:3000/employees
```

#### اضافه کردن کارمند جدید (POST)
برای اضافه کردن کارمند جدید، یک درخواست POST به `/employees` ارسال کنید و اطلاعات کارمند را به فرمت JSON ارسال کنید.

مثال:
```bash
curl -X POST http://localhost:3000/employees -d '{"name": "Ali", "email": "ali@example.com", "phone": "123456789", "position": "Developer", "salary": 5000, "age": 30, "team_id": 1}' -H "Content-Type: application/json"
```

#### به‌روزرسانی اطلاعات کارمند (PUT)
برای به‌روزرسانی اطلاعات یک کارمند، یک درخواست PUT به `/employees/:id` ارسال کنید.

مثال:
```bash
curl -X PUT http://localhost:3000/employees/1 -d '{"name": "Ali Updated", "email": "ali_updated@example.com", "phone": "987654321", "position": "Senior Developer", "salary": 6000, "age": 32, "team_id": 2}' -H "Content-Type: application/json"
```

#### حذف کارمند (DELETE)
برای حذف یک کارمند، یک درخواست DELETE به `/employees/:id` ارسال کنید.

مثال:
```bash
curl -X DELETE http://localhost:3000/employees/1
```

### 6. استفاده از رابط کاربری Vue.js

برای تعامل با API از رابط کاربری Vue.js استفاده کنید. این رابط به صورت خودکار درخواست‌ها را به سرور Go ارسال می‌کند و داده‌ها را نمایش می‌دهد.

برای مشاهده رابط کاربری، به آدرس داده شده توسط اپ برید برای مثال:

```
http://localhost:5174/
```
## Postman

### 1. **Add Employee**  
**Endpoint:** `POST /employees`  
**شرح:** این درخواست یک کارمند جدید به سیستم اضافه می‌کند. پارامترهای ورودی شامل نام، ایمیل، تلفن، سمت، حقوق، سن و شناسه تیم است.  
**Body Example:**  
```json
{
    "name": " ",
    "email": " ",
    "phone": " ",
    "position": " ",
    "salary": 0,
    "age": 0,
    "team_id": 0
}
```

---

### 2. **Update Employee**  
**Endpoint:** `PUT /employees/{id}`  
**شرح:** این درخواست اطلاعات یک کارمند را بر اساس شناسه‌ی آن به‌روزرسانی می‌کند. پارامترهای ورودی شامل نام، ایمیل، تلفن، سمت، حقوق، سن و شناسه تیم است.  
**Body Example:**  
```json
{
    "name": "",
    "email": "",
    "phone": "",
    "position": "",
    "salary": 0,
    "age": 0,
    "team_id": 0
}
```

---

### 3. **Delete Employee**  
**Endpoint:** `DELETE /employees/{id}`  
**شرح:** این درخواست یک کارمند را از سیستم حذف می‌کند بر اساس شناسه‌ی آن.  
**Body:** (بدون بدنه)

---

### 4. **Select All Employees**  
**Endpoint:** `GET /employees`  
**شرح:** این درخواست تمام اطلاعات کارمندان موجود را بازمی‌گرداند.  
**Body:** (بدون بدنه)

---

### 5. **Add Team**  
**Endpoint:** `POST /teams`  
**شرح:** این درخواست یک تیم جدید به سیستم اضافه می‌کند. پارامترهای ورودی شامل نام، توضیحات و شناسه رهبر تیم است.  
**Body Example:**  
```json
{
    "name": " ",
    "description": " ",
    "leader_id": 0
}
```

---

### 6. **Update Team**  
**Endpoint:** `PUT /teams/{id}`  
**شرح:** این درخواست اطلاعات یک تیم را بر اساس شناسه آن به‌روزرسانی می‌کند. پارامترهای ورودی شامل نام، توضیحات و شناسه رهبر تیم است.  
**Body Example:**  
```json
{
    "name": " ",
    "description": " ",
    "leader_id": 0
}
```

---

### 7. **Delete Team**  
**Endpoint:** `DELETE /teams/{id}`  
**شرح:** این درخواست یک تیم را از سیستم حذف می‌کند بر اساس شناسه آن.  
**Body:** (بدون بدنه)

---

### 8. **Select All Teams**  
**Endpoint:** `GET /teams`  
**شرح:** این درخواست تمام اطلاعات تیم‌های موجود را بازمی‌گرداند.  
**Body:** (بدون بدنه)

---


#### ساخته شده با:

- [Vue](https://vuejs.org)
- [Go-Lang](https://go.dev/)
</div>
