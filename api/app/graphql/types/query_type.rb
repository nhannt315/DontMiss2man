module Types
  class QueryType < Types::BaseObject
    include GraphQL::Types::Relay::HasNodeField
    include GraphQL::Types::Relay::HasNodesField

    field :viewer, Types::User, null: false, description: "ログイン中のユーザー"
    def viewer
      context.current_user
    end
  end
end
