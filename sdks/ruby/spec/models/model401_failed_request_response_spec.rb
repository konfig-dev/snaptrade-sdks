=begin
#SnapTrade

#Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for SnapTrade::Model401FailedRequestResponse
describe SnapTrade::Model401FailedRequestResponse do
  let(:instance) { SnapTrade::Model401FailedRequestResponse.new }

  describe 'test an instance of Model401FailedRequestResponse' do
    it 'should create an instance of Model401FailedRequestResponse' do
      expect(instance).to be_instance_of(SnapTrade::Model401FailedRequestResponse)
    end
  end
  describe 'test attribute "default_detail"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "default_code"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
