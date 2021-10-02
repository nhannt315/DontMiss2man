# frozen_string_literal: true

module Types
  class UserType < Types::BaseObject
    self.authorization_required = false

    global_id_field :id

    field :email, String, null: false, description: "ユーザーのEmail"

    def email
      object.email
    end
  end
end
