module DontMiss2Man
  module JsonFormatter
    def call(severity, timestamp, progname, msg)
      super(severity, timestamp, progname, JSON.dump(msg))
    end
  end

  class << self
    attr_accessor :debugging_logger

    attr_accessor :scraping_logger

    attr_accessor :access_logger

    def log_path
      Rails.root.join('log')
    end
  end
end

DontMiss2Man.debugging_logger = begin
                                  path = File.join(DontMiss2Man.log_path, 'real_estate_debugging.log')
                                  unless File.exist?(File.dirname(path))
                                    FileUtils.mkdir_p(File.dirname(path))
                                  end

                                  f = File.open(path, 'a')
                                  f.binmode
                                  f.sync = Rails.application.config.autoflush_log

                                  logger = ActiveSupport::Logger.new(f)
                                  logger.formatter = Logger::Formatter.new
                                  logger = ActiveSupport::TaggedLogging.new(logger)
                                  logger.formatter.extend DontMiss2Man::JsonFormatter

                                  logger
                                end

DontMiss2Man.scraping_logger = begin
                                 path = File.join(DontMiss2Man.log_path, 'real_estate_scraping.log')
                                 unless File.exist?(File.dirname(path))
                                   FileUtils.mkdir_p(File.dirname(path))
                                 end

                                 f = File.open(path, 'a')
                                 f.binmode
                                 f.sync = Rails.application.config.autoflush_log

                                 logger = ActiveSupport::Logger.new(f)
                                 logger.formatter = Logger::Formatter.new
                                 logger = ActiveSupport::TaggedLogging.new(logger)
                                 logger.formatter.extend DontMiss2Man::JsonFormatter

                                 logger
                               end

DontMiss2Man.access_logger = begin
                               path = File.join(DontMiss2Man.log_path, 'real_estate_access.log')
                               unless File.exist?(File.dirname(path))
                                 FileUtils.mkdir_p(File.dirname(path))
                               end

                               f = File.open(path, 'a')
                               f.binmode
                               f.sync = Rails.application.config.autoflush_log

                               logger = ActiveSupport::Logger.new(f)
                               logger.formatter = Logger::Formatter.new
                               logger = ActiveSupport::TaggedLogging.new(logger)
                               logger.formatter.extend DontMiss2Man::JsonFormatter

                               logger
                             end
