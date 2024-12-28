// Types for error handling
type ErrorResponse = {
  message: string;
  status?: number;
};

// Convert any error to a structured error response
const normalizeError = (error: any): ErrorResponse => {
  if (error.response) {
    return {
      message: error.response.data?.message || 'خطای سرور',
      status: error.response.status
    };
  }
  
  if (error instanceof Error) {
    return {
      message: error.message
    };
  }
  
  return {
    message: 'خطای ناشناخته'
  };
};

// Safe error logging that avoids Symbol serialization issues
const logError = (context: string, error: any) => {
  const errorInfo = {
    context,
    message: error?.message || 'Unknown error',
    status: error?.response?.status,
    data: error?.response?.data
  };
  
  console.error('Application Error:', JSON.stringify(errorInfo, null, 2));
};

export const handleApiError = (error: any, context: string): string => {
  // Safely log the error
  logError(context, error);
  
  // Get normalized error response
  const errorResponse = normalizeError(error);
  
  // Return appropriate message based on error type
  if (errorResponse.status) {
    switch (errorResponse.status) {
      case 400:
        return 'داده‌های ورودی نامعتبر هستند';
      case 404:
        return 'اطلاعات مورد نظر یافت نشد';
      case 500:
        return 'خطای سرور رخ داده است';
      default:
        return 'خطایی در عملیات رخ داده است';
    }
  }
  
  return errorResponse.message || 'خطا در برقراری ارتباط با سرور';
};