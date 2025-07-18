export const validateEmail = (email) => {
  const emailRegex = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return emailRegex.test(email) === true;
};

export const validatePassword = (password) => {
  const lengthCheck = password.length >= 8;
  const uppercaseCheck = /[A-Z]/.test(password);
  const lowercaseCheck = /[a-z]/.test(password);
  const numberCheck = /\d/.test(password);
  const specialCharCheck = /[@$!%*?&]/.test(password);
  let errors = []

  if (!lengthCheck) errors.push("Password must be at least 8 characters long.");
  if (!uppercaseCheck) errors.push("Password must contain at least one uppercase letter.");
  if (!lowercaseCheck) errors.push("Password must contain at least one lowercase letter.");
  if (!numberCheck) errors.push("Password must contain at least one number.");
  if (!specialCharCheck) errors.push("Password must contain at least one special character (@, $, !, %, *, ?, &).");

  return errors; // Password is valid
};


export const getErrorMessageFromResponse = (error) => {
  return (error.response?.data?.error || error.message || error.response);
}

export function formatNumber(num) {
  if (num < 1000) {
    return num.toString(); // Return as is for numbers less than 1000
  } else if (num < 1_000_000) {
    return (num / 1000).toFixed(1).replace(/\.0$/, "") + "k"; // Convert to 'k'
  } else if (num < 1_000_000_000) {
    return (num / 1_000_000).toFixed(1).replace(/\.0$/, "") + "M"; // Convert to 'M'
  } else {
    return (num / 1_000_000_000).toFixed(1).replace(/\.0$/, "") + "B"; // Convert to 'B'
  }
}

export function generateOAuthState() {
  if (window.crypto && window.crypto.randomUUID) {
    return window.crypto.randomUUID();
  } else {
    // Fallback for older browsers (less secure, consider a polyfill for production)
    console.warn("crypto.randomUUID() not available. Using less secure fallback for state generation.");
    const array = new Uint32Array(28);
    window.crypto.getRandomValues(array);
    return Array.from(array, dec => ('0' + dec.toString(16)).slice(-2)).join('');
  }
}

export function capitalize(str) {
  if (!str) return "";
  return str.charAt(0).toUpperCase() + str.slice(1);
}