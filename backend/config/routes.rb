Rails.application.routes.draw do
  mount_devise_token_auth_for "User", at: "api/v1/auth"
  namespace :api do
    namespace :v1 do
      resources :buildings
      resources :rooms
      resources :favorites do
        collection do
          post "create"
          post "delete"
        end
      end
    end
  end
end
