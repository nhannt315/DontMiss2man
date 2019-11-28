Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      resources :buildings
      resources :rooms
    end
  end
end
