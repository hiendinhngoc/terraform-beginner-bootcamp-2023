variable "user_uuid" {
  type = string
  validation {
    condition = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
    error_message = "User UUID must be in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  }
}

variable "bucket_name" {
  type = string

  validation {
    condition = (
      length(var.bucket_name) >= 3 && length(var.bucket_name) <= 63 &&
      can(regex("^([a-z0-9]{1}[a-z0-9-]{1,61}[a-z0-9]{1})$", var.bucket_name))
    )
    error_message = "Bucket name must be between 3 and 63 characters long. Invalid bucket name, please check https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html for more details."
  }
}

variable "index_html_filepath" {
  description = "Path to the index.html file"
  type        = string
  validation {
    condition     = fileexists(var.index_html_filepath)
    error_message = "The specified index.html file does not exist."
  }
}

variable "error_html_filepath" {
  description = "Path to the error.html file"
  type        = string
  validation {
    condition     = fileexists(var.error_html_filepath)
    error_message = "The specified error.html file does not exist."
  }
}