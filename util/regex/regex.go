package regex

// PhoneRegex regex
const PhoneRegex = `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`

// UsernameRegex regex
const UsernameRegex = `^[a-zA-Z0-9]+$`

// SlugRegex regex
const SlugRegex = `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`
