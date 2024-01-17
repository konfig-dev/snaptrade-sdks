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
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.time.LocalDate;
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
 * SubPeriodReturnRate
 */@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class SubPeriodReturnRate {
  public static final String SERIALIZED_NAME_PERIOD_START = "periodStart";
  @SerializedName(SERIALIZED_NAME_PERIOD_START)
  private LocalDate periodStart;

  public static final String SERIALIZED_NAME_PERIOD_END = "periodEnd";
  @SerializedName(SERIALIZED_NAME_PERIOD_END)
  private LocalDate periodEnd;

  public static final String SERIALIZED_NAME_RATE_OF_RETURN = "rateOfReturn";
  @SerializedName(SERIALIZED_NAME_RATE_OF_RETURN)
  private Double rateOfReturn;

  public SubPeriodReturnRate() {
  }

  public SubPeriodReturnRate periodStart(LocalDate periodStart) {
    
    
    
    
    this.periodStart = periodStart;
    return this;
  }

   /**
   * Date used to specify timeframe for a reporting call (in YYYY-MM-DD format)
   * @return periodStart
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Mon Jan 24 00:00:00 UTC 2022", value = "Date used to specify timeframe for a reporting call (in YYYY-MM-DD format)")

  public LocalDate getPeriodStart() {
    return periodStart;
  }


  public void setPeriodStart(LocalDate periodStart) {
    
    
    
    this.periodStart = periodStart;
  }


  public SubPeriodReturnRate periodEnd(LocalDate periodEnd) {
    
    
    
    
    this.periodEnd = periodEnd;
    return this;
  }

   /**
   * Date used to specify timeframe for a reporting call (in YYYY-MM-DD format)
   * @return periodEnd
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Mon Jan 24 00:00:00 UTC 2022", value = "Date used to specify timeframe for a reporting call (in YYYY-MM-DD format)")

  public LocalDate getPeriodEnd() {
    return periodEnd;
  }


  public void setPeriodEnd(LocalDate periodEnd) {
    
    
    
    this.periodEnd = periodEnd;
  }


  public SubPeriodReturnRate rateOfReturn(Double rateOfReturn) {
    
    
    
    
    this.rateOfReturn = rateOfReturn;
    return this;
  }

  public SubPeriodReturnRate rateOfReturn(Integer rateOfReturn) {
    
    
    
    
    this.rateOfReturn = rateOfReturn.doubleValue();
    return this;
  }

   /**
   * The return rate for the given period
   * @return rateOfReturn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "0.012312367452", value = "The return rate for the given period")

  public Double getRateOfReturn() {
    return rateOfReturn;
  }


  public void setRateOfReturn(Double rateOfReturn) {
    
    
    
    this.rateOfReturn = rateOfReturn;
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
   * @return the SubPeriodReturnRate instance itself
   */
  public SubPeriodReturnRate putAdditionalProperty(String key, Object value) {
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
    SubPeriodReturnRate subPeriodReturnRate = (SubPeriodReturnRate) o;
    return Objects.equals(this.periodStart, subPeriodReturnRate.periodStart) &&
        Objects.equals(this.periodEnd, subPeriodReturnRate.periodEnd) &&
        Objects.equals(this.rateOfReturn, subPeriodReturnRate.rateOfReturn)&&
        Objects.equals(this.additionalProperties, subPeriodReturnRate.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(periodStart, periodEnd, rateOfReturn, additionalProperties);
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
    sb.append("class SubPeriodReturnRate {\n");
    sb.append("    periodStart: ").append(toIndentedString(periodStart)).append("\n");
    sb.append("    periodEnd: ").append(toIndentedString(periodEnd)).append("\n");
    sb.append("    rateOfReturn: ").append(toIndentedString(rateOfReturn)).append("\n");
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
    openapiFields.add("periodStart");
    openapiFields.add("periodEnd");
    openapiFields.add("rateOfReturn");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to SubPeriodReturnRate
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!SubPeriodReturnRate.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in SubPeriodReturnRate is not found in the empty JSON string", SubPeriodReturnRate.openapiRequiredFields.toString()));
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!SubPeriodReturnRate.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'SubPeriodReturnRate' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<SubPeriodReturnRate> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(SubPeriodReturnRate.class));

       return (TypeAdapter<T>) new TypeAdapter<SubPeriodReturnRate>() {
           @Override
           public void write(JsonWriter out, SubPeriodReturnRate value) throws IOException {
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
           public SubPeriodReturnRate read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             SubPeriodReturnRate instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of SubPeriodReturnRate given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of SubPeriodReturnRate
  * @throws IOException if the JSON string is invalid with respect to SubPeriodReturnRate
  */
  public static SubPeriodReturnRate fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, SubPeriodReturnRate.class);
  }

 /**
  * Convert an instance of SubPeriodReturnRate to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

