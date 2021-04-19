#!/usr/bin/env rake

ENV['APP_ENV'] ||= 'development'

require 'dotenv'
require 'standalone_migrations'
require 'lhm'

if ENV['APP_ENV'] == 'development'
  Dotenv.load
end

StandaloneMigrations::Tasks.load_tasks

namespace :lhm do
  desc 'Run Lhm.cleanup'
  task :cleanup do
    ActiveRecord::Base.establish_connection
    Lhm.cleanup(true)
    Rake::Task['db:schema:dump'].invoke
  end
end