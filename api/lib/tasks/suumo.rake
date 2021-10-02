namespace :suumo do
    desc "Crawl building and rooms info from suumo"
    task crawl: :environment do
        Office.all.each do |office|
            Suumo::Scraper.new(office).execute!
          end
    end
end