json.data do
  json.list @building_list do |building|
    json.merge! building.attributes
    json.rooms building.rooms do |room|
      json.merge! room.attributes
    end
  end
  json.total_pages @building_list.total_pages
  json.total @building_list.total_count
  json.page @page
end
