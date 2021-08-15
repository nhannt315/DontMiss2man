module RequestSpecHelper
  def json
    JSON.parse(response.body)
  end

  # return valid headers
  def valid_headers
    {
        :Accept => "application/json",
        "Content-Type" => "application/json"
    }
  end

end
