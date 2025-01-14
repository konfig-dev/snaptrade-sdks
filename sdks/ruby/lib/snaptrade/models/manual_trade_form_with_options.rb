=begin
#SnapTrade

#Connect brokerage accounts to your app for live positions and trading

The version of the OpenAPI document: 1.0.0
Contact: api@snaptrade.com
=end

require 'date'
require 'time'

module SnapTrade
  # Inputs for placing an order with the brokerage.
  class ManualTradeFormWithOptions
    # Unique identifier for the connected brokerage account. This is the UUID used to reference the account in SnapTrade.
    attr_accessor :account_id

    # The action describes the intent or side of a trade. This is either `BUY` or `SELL` for Equity symbols or `BUY_TO_OPEN`, `BUY_TO_CLOSE`, `SELL_TO_OPEN` or `SELL_TO_CLOSE` for Options.
    attr_accessor :action

    # The universal symbol ID of the security to trade. Must be 'null' if `symbol` is provided, otherwise must be provided.
    attr_accessor :universal_symbol_id

    # The security's trading ticker symbol. This currently only support Options symbols in the 21 character OCC format. For example \"AAPL  131124C00240000\" represents a call option on AAPL expiring on 2024-11-13 with a strike price of $240. For more information on the OCC format, see [here](https://en.wikipedia.org/wiki/Option_symbol#OCC_format). If 'symbol' is provided, then 'universal_symbol_id' must be 'null'.
    attr_accessor :symbol

    # The type of order to place.  - For `Limit` and `StopLimit` orders, the `price` field is required. - For `Stop` and `StopLimit` orders, the `stop` field is required. 
    attr_accessor :order_type

    # The Time in Force type for the order. This field indicates how long the order will remain active before it is executed or expires. Here are the supported values:   - `Day` - Day. The order is valid only for the trading day on which it is placed.   - `GTC` - Good Til Canceled. The order is valid until it is executed or canceled.   - `FOK` - Fill Or Kill. The order must be executed in its entirety immediately or be canceled completely. 
    attr_accessor :time_in_force

    # The limit price for `Limit` and `StopLimit` orders.
    attr_accessor :price

    # The price at which a stop order is triggered for `Stop` and `StopLimit` orders.
    attr_accessor :stop

    # For Equity orders, this represents the number of shares for the order. This can be a decimal for fractional orders. Must be `null` if `notional_value` is provided. If placing an Option order, this field represents the number of contracts to buy or sell. (e.g., 1 contract = 100 shares).
    attr_accessor :units

    attr_accessor :notional_value

    # The class of order intended to be placed. Defaults to SIMPLE for regular, one legged trades. Set to BRACKET if looking to place a bracket (One-triggers-a-one-cancels-the-other) order, then specify take profit and stop loss conditions. Bracket orders currently only supported on Alpaca, Tradier, and Tradestation, contact us for more details
    attr_accessor :order_class

    attr_accessor :stop_loss

    attr_accessor :take_profit

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'account_id' => :'account_id',
        :'action' => :'action',
        :'universal_symbol_id' => :'universal_symbol_id',
        :'symbol' => :'symbol',
        :'order_type' => :'order_type',
        :'time_in_force' => :'time_in_force',
        :'price' => :'price',
        :'stop' => :'stop',
        :'units' => :'units',
        :'notional_value' => :'notional_value',
        :'order_class' => :'order_class',
        :'stop_loss' => :'stop_loss',
        :'take_profit' => :'take_profit'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'account_id' => :'String',
        :'action' => :'ActionStrictWithOptions',
        :'universal_symbol_id' => :'String',
        :'symbol' => :'String',
        :'order_type' => :'OrderTypeStrict',
        :'time_in_force' => :'TimeInForceStrict',
        :'price' => :'Float',
        :'stop' => :'Float',
        :'units' => :'Float',
        :'notional_value' => :'ManualTradeFormNotionalValue',
        :'order_class' => :'OrderClass',
        :'stop_loss' => :'ManualTradeFormWithOptionsStopLoss',
        :'take_profit' => :'ManualTradeFormWithOptionsTakeProfit'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'universal_symbol_id',
        :'symbol',
        :'price',
        :'stop',
        :'units',
        :'notional_value',
        :'order_class',
        :'stop_loss',
        :'take_profit'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `SnapTrade::ManualTradeFormWithOptions` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `SnapTrade::ManualTradeFormWithOptions`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'account_id')
        self.account_id = attributes[:'account_id']
      end

      if attributes.key?(:'action')
        self.action = attributes[:'action']
      end

      if attributes.key?(:'universal_symbol_id')
        self.universal_symbol_id = attributes[:'universal_symbol_id']
      end

      if attributes.key?(:'symbol')
        self.symbol = attributes[:'symbol']
      end

      if attributes.key?(:'order_type')
        self.order_type = attributes[:'order_type']
      end

      if attributes.key?(:'time_in_force')
        self.time_in_force = attributes[:'time_in_force']
      end

      if attributes.key?(:'price')
        self.price = attributes[:'price']
      end

      if attributes.key?(:'stop')
        self.stop = attributes[:'stop']
      end

      if attributes.key?(:'units')
        self.units = attributes[:'units']
      end

      if attributes.key?(:'notional_value')
        self.notional_value = attributes[:'notional_value']
      end

      if attributes.key?(:'order_class')
        self.order_class = attributes[:'order_class']
      end

      if attributes.key?(:'stop_loss')
        self.stop_loss = attributes[:'stop_loss']
      end

      if attributes.key?(:'take_profit')
        self.take_profit = attributes[:'take_profit']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if @account_id.nil?
        invalid_properties.push('invalid value for "account_id", account_id cannot be nil.')
      end

      if @action.nil?
        invalid_properties.push('invalid value for "action", action cannot be nil.')
      end

      if @order_type.nil?
        invalid_properties.push('invalid value for "order_type", order_type cannot be nil.')
      end

      if @time_in_force.nil?
        invalid_properties.push('invalid value for "time_in_force", time_in_force cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @account_id.nil?
      return false if @action.nil?
      return false if @order_type.nil?
      return false if @time_in_force.nil?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          account_id == o.account_id &&
          action == o.action &&
          universal_symbol_id == o.universal_symbol_id &&
          symbol == o.symbol &&
          order_type == o.order_type &&
          time_in_force == o.time_in_force &&
          price == o.price &&
          stop == o.stop &&
          units == o.units &&
          notional_value == o.notional_value &&
          order_class == o.order_class &&
          stop_loss == o.stop_loss &&
          take_profit == o.take_profit
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [account_id, action, universal_symbol_id, symbol, order_type, time_in_force, price, stop, units, notional_value, order_class, stop_loss, take_profit].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      new.build_from_hash(attributes)
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      attributes = attributes.transform_keys(&:to_sym)
      self.class.openapi_types.each_pair do |key, type|
        if attributes[self.class.attribute_map[key]].nil? && self.class.openapi_nullable.include?(key)
          self.send("#{key}=", nil)
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[self.class.attribute_map[key]].is_a?(Array)
            self.send("#{key}=", attributes[self.class.attribute_map[key]].map { |v| _deserialize($1, v) })
          end
        elsif !attributes[self.class.attribute_map[key]].nil?
          self.send("#{key}=", _deserialize(type, attributes[self.class.attribute_map[key]]))
        end
      end

      self
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def _deserialize(type, value)
      case type.to_sym
      when :Time
        Time.parse(value)
      when :Date
        Date.parse(value)
      when :String
        value.to_s
      when :Integer
        value.to_i
      when :Float
        value.to_f
      when :Boolean
        if value.to_s =~ /\A(true|t|yes|y|1)\z/i
          true
        else
          false
        end
      when :Object
        # generic object (usually a Hash), return directly
        value
      when /\AArray<(?<inner_type>.+)>\z/
        inner_type = Regexp.last_match[:inner_type]
        value.map { |v| _deserialize(inner_type, v) }
      when /\AHash<(?<k_type>.+?), (?<v_type>.+)>\z/
        k_type = Regexp.last_match[:k_type]
        v_type = Regexp.last_match[:v_type]
        {}.tap do |hash|
          value.each do |k, v|
            hash[_deserialize(k_type, k)] = _deserialize(v_type, v)
          end
        end
      else # model
        # models (e.g. Pet) or oneOf
        klass = SnapTrade.const_get(type)
        klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
      end
    end

    # Returns the string representation of the object
    # @return [String] String presentation of the object
    def to_s
      to_hash.to_s
    end

    # to_body is an alias to to_hash (backward compatibility)
    # @return [Hash] Returns the object in the form of hash
    def to_body
      to_hash
    end

    # Returns the object in the form of hash
    # @return [Hash] Returns the object in the form of hash
    def to_hash
      hash = {}
      self.class.attribute_map.each_pair do |attr, param|
        value = self.send(attr)
        if value.nil?
          is_nullable = self.class.openapi_nullable.include?(attr)
          next if !is_nullable || (is_nullable && !instance_variable_defined?(:"@#{attr}"))
        end

        hash[param] = _to_hash(value)
      end
      hash
    end

    # Outputs non-array value in the form of hash
    # For object, use to_hash. Otherwise, just return the value
    # @param [Object] value Any valid value
    # @return [Hash] Returns the value in the form of hash
    def _to_hash(value)
      if value.is_a?(Array)
        value.compact.map { |v| _to_hash(v) }
      elsif value.is_a?(Hash)
        {}.tap do |hash|
          value.each { |k, v| hash[k] = _to_hash(v) }
        end
      elsif value.respond_to? :to_hash
        value.to_hash
      else
        value
      end
    end

  end

end
