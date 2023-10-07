/*
 * SnapTrade
 * Connect brokerage accounts to your app for live positions and trading
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: api@snaptrade.com
 *
 * NOTE: This class is auto generated by Konfig (https://konfigthis.com).
 * Do not edit the class manually.
 */


package com.konfigthis.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import com.konfigthis.client.model.Currency;
import com.konfigthis.client.model.Exchange;
import com.konfigthis.client.model.OptionStrategy;
import com.konfigthis.client.model.OptionStrategyLegsInner;
import com.konfigthis.client.model.SecurityType;
import com.konfigthis.client.model.StrategyQuotesGreek;
import com.konfigthis.client.model.UniversalSymbol;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.UUID;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import org.apache.commons.lang3.StringUtils;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import com.konfigthis.client.JSON;

/**
 * StrategyQuotes
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class StrategyQuotes {
  public static final String SERIALIZED_NAME_STRATEGY = "strategy";
  @SerializedName(SERIALIZED_NAME_STRATEGY)
  private OptionStrategy strategy;

  public static final String SERIALIZED_NAME_OPEN_PRICE = "open_price";
  @SerializedName(SERIALIZED_NAME_OPEN_PRICE)
  private Double openPrice;

  public static final String SERIALIZED_NAME_BID_PRICE = "bid_price";
  @SerializedName(SERIALIZED_NAME_BID_PRICE)
  private Double bidPrice;

  public static final String SERIALIZED_NAME_ASK_PRICE = "ask_price";
  @SerializedName(SERIALIZED_NAME_ASK_PRICE)
  private Double askPrice;

  public static final String SERIALIZED_NAME_VOLATILITY = "volatility";
  @SerializedName(SERIALIZED_NAME_VOLATILITY)
  private Double volatility;

  public static final String SERIALIZED_NAME_GREEK = "greek";
  @SerializedName(SERIALIZED_NAME_GREEK)
  private StrategyQuotesGreek greek;

  public StrategyQuotes() {
  }

  public StrategyQuotes strategy(OptionStrategy strategy) {
    
    
    
    
    this.strategy = strategy;
    return this;
  }

   /**
   * Get strategy
   * @return strategy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OptionStrategy getStrategy() {
    return strategy;
  }


  public void setStrategy(OptionStrategy strategy) {
    
    
    
    this.strategy = strategy;
  }


  public StrategyQuotes openPrice(Double openPrice) {
    
    
    
    
    this.openPrice = openPrice;
    return this;
  }

  public StrategyQuotes openPrice(Integer openPrice) {
    
    
    
    
    this.openPrice = openPrice.doubleValue();
    return this;
  }

   /**
   * Trade Price if limit or stop limit order
   * @return openPrice
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "31.33", value = "Trade Price if limit or stop limit order")

  public Double getOpenPrice() {
    return openPrice;
  }


  public void setOpenPrice(Double openPrice) {
    
    
    
    this.openPrice = openPrice;
  }


  public StrategyQuotes bidPrice(Double bidPrice) {
    
    
    
    
    this.bidPrice = bidPrice;
    return this;
  }

  public StrategyQuotes bidPrice(Integer bidPrice) {
    
    
    
    
    this.bidPrice = bidPrice.doubleValue();
    return this;
  }

   /**
   * Trade Price if limit or stop limit order
   * @return bidPrice
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "31.33", value = "Trade Price if limit or stop limit order")

  public Double getBidPrice() {
    return bidPrice;
  }


  public void setBidPrice(Double bidPrice) {
    
    
    
    this.bidPrice = bidPrice;
  }


  public StrategyQuotes askPrice(Double askPrice) {
    
    
    
    
    this.askPrice = askPrice;
    return this;
  }

  public StrategyQuotes askPrice(Integer askPrice) {
    
    
    
    
    this.askPrice = askPrice.doubleValue();
    return this;
  }

   /**
   * Trade Price if limit or stop limit order
   * @return askPrice
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "31.33", value = "Trade Price if limit or stop limit order")

  public Double getAskPrice() {
    return askPrice;
  }


  public void setAskPrice(Double askPrice) {
    
    
    
    this.askPrice = askPrice;
  }


  public StrategyQuotes volatility(Double volatility) {
    
    
    
    
    this.volatility = volatility;
    return this;
  }

  public StrategyQuotes volatility(Integer volatility) {
    
    
    
    
    this.volatility = volatility.doubleValue();
    return this;
  }

   /**
   * Get volatility
   * @return volatility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "0.141", value = "")

  public Double getVolatility() {
    return volatility;
  }


  public void setVolatility(Double volatility) {
    
    
    
    this.volatility = volatility;
  }


  public StrategyQuotes greek(StrategyQuotesGreek greek) {
    
    
    
    
    this.greek = greek;
    return this;
  }

   /**
   * Get greek
   * @return greek
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public StrategyQuotesGreek getGreek() {
    return greek;
  }


  public void setGreek(StrategyQuotesGreek greek) {
    
    
    
    this.greek = greek;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the StrategyQuotes instance itself
   */
  public StrategyQuotes putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    StrategyQuotes strategyQuotes = (StrategyQuotes) o;
    return Objects.equals(this.strategy, strategyQuotes.strategy) &&
        Objects.equals(this.openPrice, strategyQuotes.openPrice) &&
        Objects.equals(this.bidPrice, strategyQuotes.bidPrice) &&
        Objects.equals(this.askPrice, strategyQuotes.askPrice) &&
        Objects.equals(this.volatility, strategyQuotes.volatility) &&
        Objects.equals(this.greek, strategyQuotes.greek)&&
        Objects.equals(this.additionalProperties, strategyQuotes.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(strategy, openPrice, bidPrice, askPrice, volatility, greek, additionalProperties);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class StrategyQuotes {\n");
    sb.append("    strategy: ").append(toIndentedString(strategy)).append("\n");
    sb.append("    openPrice: ").append(toIndentedString(openPrice)).append("\n");
    sb.append("    bidPrice: ").append(toIndentedString(bidPrice)).append("\n");
    sb.append("    askPrice: ").append(toIndentedString(askPrice)).append("\n");
    sb.append("    volatility: ").append(toIndentedString(volatility)).append("\n");
    sb.append("    greek: ").append(toIndentedString(greek)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("strategy");
    openapiFields.add("open_price");
    openapiFields.add("bid_price");
    openapiFields.add("ask_price");
    openapiFields.add("volatility");
    openapiFields.add("greek");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to StrategyQuotes
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!StrategyQuotes.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in StrategyQuotes is not found in the empty JSON string", StrategyQuotes.openapiRequiredFields.toString()));
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!StrategyQuotes.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'StrategyQuotes' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<StrategyQuotes> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(StrategyQuotes.class));

       return (TypeAdapter<T>) new TypeAdapter<StrategyQuotes>() {
           @Override
           public void write(JsonWriter out, StrategyQuotes value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additonal properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public StrategyQuotes read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             StrategyQuotes instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of StrategyQuotes given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of StrategyQuotes
  * @throws IOException if the JSON string is invalid with respect to StrategyQuotes
  */
  public static StrategyQuotes fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, StrategyQuotes.class);
  }

 /**
  * Convert an instance of StrategyQuotes to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

