=begin
#SnapTrade

#Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for SnapTrade::UserIDandSecret
describe SnapTrade::UserIDandSecret do
  let(:instance) { SnapTrade::UserIDandSecret.new }

  describe 'test an instance of UserIDandSecret' do
    it 'should create an instance of UserIDandSecret' do
      expect(instance).to be_instance_of(SnapTrade::UserIDandSecret)
    end
  end
  describe 'test attribute "user_id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "user_secret"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
