json.data do
  json.list @building_list do |building|
    json.merge! building.attributes
    json.rooms building.rooms
  end
  json.total_pages @building_list.total_pages
  json.page @page
end
