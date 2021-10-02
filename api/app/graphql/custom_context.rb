# frozen_string_literal: true

class CustomContext < GraphQL::Query::Context
  def current_user
    self[:current_user]
  end

  def cache(key)
    self[key] ||= yield
  end
end
