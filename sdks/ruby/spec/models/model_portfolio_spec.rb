=begin
#SnapTrade

#Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for SnapTrade::ModelPortfolio
describe SnapTrade::ModelPortfolio do
  let(:instance) { SnapTrade::ModelPortfolio.new }

  describe 'test an instance of ModelPortfolio' do
    it 'should create an instance of ModelPortfolio' do
      expect(instance).to be_instance_of(SnapTrade::ModelPortfolio)
    end
  end
  describe 'test attribute "id"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "name"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "model_type"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('Integer', [-1, 0, 1])
      # validator.allowable_values.each do |value|
      #   expect { instance.model_type = value }.not_to raise_error
      # end
    end
  end

end
