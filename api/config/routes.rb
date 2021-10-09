Rails.application.routes.draw do
  post "/graphql", to: "graphql#execute"
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

      post "/auth/login", to: "authentications#login"
      post "/auth/register", to: "authentications#register"

      resources :users do
        collection do
          get "info"
        end
      end
    end
  end
end
