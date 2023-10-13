# Terraform Beginner Bootcamp 2023 - Week 2

## Working with Ruby

### Bundler

Bundler is a package manager for Ruby. It's the primary way to install Ruby packages (known as gems) for Ruby projects.

#### Gems

You need to create a Gemfile and define your gems in that file.
```ruby
source "https://rubygems.org"

gem 'sinatra'
gem 'rake'
gem 'pry'
gem 'puma'
gem 'activerecord'
```

then you need to run the `bundle install` (or `bundle` for short)

This will install the gems on the system globally (unlike nodejs which install packages in place in a folder called node_modules)

A Gemfile.lock will be created to lock down the gem versions used in this project.

#### Executing Ruby scripts in the context of bundler

We have to use `bundle exec` to tell future Ruby scripts to use the gems we installed. This is the way we set context.

### Sinatra

It's a micro web-framework for Ruby to build web-apps.

Its great for mock or development servers or for very simple projects.

You can create a web-server in a single file.

https://sinatrarb.com/

## Terratowns Mock Server

### Running the web server

We can run the web server by executing the following command:
```ruby
bundle install
bundle exec ruby server.rb
```