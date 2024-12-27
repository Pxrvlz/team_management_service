// Debug levels
export const DEBUG_LEVELS = {
  INFO: 'INFO',
  WARNING: 'WARNING',
  ERROR: 'ERROR'
} as const;

type DebugLevel = typeof DEBUG_LEVELS[keyof typeof DEBUG_LEVELS];

// Debug configuration
const DEBUG_CONFIG = {
  enabled: true,
  logToConsole: true,
  showNotifications: true
};

declare global {
  interface Window {
    debugNotify?: (notification: {
      level: DebugLevel;
      context: string;
      message: string;
      timestamp: string;
      data?: any;
    }) => void;
  }
}

class DebugLogger {
  private static log(level: DebugLevel, context: string, message: string, data?: any) {
    if (!DEBUG_CONFIG.enabled) return;

    const timestamp = new Date().toISOString();
    const logMessage = {
      timestamp,
      level,
      context,
      message,
      data: data || null
    };

    if (DEBUG_CONFIG.logToConsole) {
      console.log(`[${timestamp}] ${level} - ${context}: ${message}`, data ? data : '');
    }

    if (DEBUG_CONFIG.showNotifications && window.debugNotify) {
      window.debugNotify(logMessage);
    }

    return logMessage;
  }

  static info(context: string, message: string, data?: any) {
    return this.log(DEBUG_LEVELS.INFO, context, message, data);
  }

  static warning(context: string, message: string, data?: any) {
    return this.log(DEBUG_LEVELS.WARNING, context, message, data);
  }

  static error(context: string, message: string, data?: any) {
    return this.log(DEBUG_LEVELS.ERROR, context, message, data);
  }
}

export default DebugLogger;