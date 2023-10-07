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
 * Stock Exchange
 */
@ApiModel(description = "Stock Exchange")@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class Exchange {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private UUID id;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_MIC_CODE = "mic_code";
  @SerializedName(SERIALIZED_NAME_MIC_CODE)
  private String micCode;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_TIMEZONE = "timezone";
  @SerializedName(SERIALIZED_NAME_TIMEZONE)
  private String timezone;

  public static final String SERIALIZED_NAME_START_TIME = "start_time";
  @SerializedName(SERIALIZED_NAME_START_TIME)
  private String startTime;

  public static final String SERIALIZED_NAME_CLOSE_TIME = "close_time";
  @SerializedName(SERIALIZED_NAME_CLOSE_TIME)
  private String closeTime;

  public static final String SERIALIZED_NAME_SUFFIX = "suffix";
  @SerializedName(SERIALIZED_NAME_SUFFIX)
  private String suffix;

  public Exchange() {
  }

  public Exchange id(UUID id) {
    
    
    
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2bcd7cc3-e922-4976-bce1-9858296801c3", value = "")

  public UUID getId() {
    return id;
  }


  public void setId(UUID id) {
    
    
    
    this.id = id;
  }


  public Exchange code(String code) {
    
    
    
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "TSX", value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    
    
    
    this.code = code;
  }


  public Exchange micCode(String micCode) {
    
    
    
    
    this.micCode = micCode;
    return this;
  }

   /**
   * Get micCode
   * @return micCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "XTSE", value = "")

  public String getMicCode() {
    return micCode;
  }


  public void setMicCode(String micCode) {
    
    
    
    this.micCode = micCode;
  }


  public Exchange name(String name) {
    
    
    
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Toronto Stock Exchange", value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    
    
    
    this.name = name;
  }


  public Exchange timezone(String timezone) {
    
    
    
    
    this.timezone = timezone;
    return this;
  }

   /**
   * Get timezone
   * @return timezone
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "America/New_York", value = "")

  public String getTimezone() {
    return timezone;
  }


  public void setTimezone(String timezone) {
    
    
    
    this.timezone = timezone;
  }


  public Exchange startTime(String startTime) {
    
    
    
    
    this.startTime = startTime;
    return this;
  }

   /**
   * Get startTime
   * @return startTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "09:30:00", value = "")

  public String getStartTime() {
    return startTime;
  }


  public void setStartTime(String startTime) {
    
    
    
    this.startTime = startTime;
  }


  public Exchange closeTime(String closeTime) {
    
    
    
    
    this.closeTime = closeTime;
    return this;
  }

   /**
   * Get closeTime
   * @return closeTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "57600", value = "")

  public String getCloseTime() {
    return closeTime;
  }


  public void setCloseTime(String closeTime) {
    
    
    
    this.closeTime = closeTime;
  }


  public Exchange suffix(String suffix) {
    
    
    
    
    this.suffix = suffix;
    return this;
  }

   /**
   * Get suffix
   * @return suffix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = ".TO", value = "")

  public String getSuffix() {
    return suffix;
  }


  public void setSuffix(String suffix) {
    
    
    
    this.suffix = suffix;
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
   * @return the Exchange instance itself
   */
  public Exchange putAdditionalProperty(String key, Object value) {
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
    Exchange exchange = (Exchange) o;
    return Objects.equals(this.id, exchange.id) &&
        Objects.equals(this.code, exchange.code) &&
        Objects.equals(this.micCode, exchange.micCode) &&
        Objects.equals(this.name, exchange.name) &&
        Objects.equals(this.timezone, exchange.timezone) &&
        Objects.equals(this.startTime, exchange.startTime) &&
        Objects.equals(this.closeTime, exchange.closeTime) &&
        Objects.equals(this.suffix, exchange.suffix)&&
        Objects.equals(this.additionalProperties, exchange.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, code, micCode, name, timezone, startTime, closeTime, suffix, additionalProperties);
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
    sb.append("class Exchange {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    micCode: ").append(toIndentedString(micCode)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    timezone: ").append(toIndentedString(timezone)).append("\n");
    sb.append("    startTime: ").append(toIndentedString(startTime)).append("\n");
    sb.append("    closeTime: ").append(toIndentedString(closeTime)).append("\n");
    sb.append("    suffix: ").append(toIndentedString(suffix)).append("\n");
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
    openapiFields.add("id");
    openapiFields.add("code");
    openapiFields.add("mic_code");
    openapiFields.add("name");
    openapiFields.add("timezone");
    openapiFields.add("start_time");
    openapiFields.add("close_time");
    openapiFields.add("suffix");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to Exchange
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!Exchange.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in Exchange is not found in the empty JSON string", Exchange.openapiRequiredFields.toString()));
        }
      }
      if ((jsonObj.get("id") != null && !jsonObj.get("id").isJsonNull()) && !jsonObj.get("id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("id").toString()));
      }
      if ((jsonObj.get("code") != null && !jsonObj.get("code").isJsonNull()) && !jsonObj.get("code").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `code` to be a primitive type in the JSON string but got `%s`", jsonObj.get("code").toString()));
      }
      if ((jsonObj.get("mic_code") != null && !jsonObj.get("mic_code").isJsonNull()) && !jsonObj.get("mic_code").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `mic_code` to be a primitive type in the JSON string but got `%s`", jsonObj.get("mic_code").toString()));
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("timezone") != null && !jsonObj.get("timezone").isJsonNull()) && !jsonObj.get("timezone").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `timezone` to be a primitive type in the JSON string but got `%s`", jsonObj.get("timezone").toString()));
      }
      if ((jsonObj.get("start_time") != null && !jsonObj.get("start_time").isJsonNull()) && !jsonObj.get("start_time").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `start_time` to be a primitive type in the JSON string but got `%s`", jsonObj.get("start_time").toString()));
      }
      if ((jsonObj.get("close_time") != null && !jsonObj.get("close_time").isJsonNull()) && !jsonObj.get("close_time").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `close_time` to be a primitive type in the JSON string but got `%s`", jsonObj.get("close_time").toString()));
      }
      if (!jsonObj.get("suffix").isJsonNull() && (jsonObj.get("suffix") != null && !jsonObj.get("suffix").isJsonNull()) && !jsonObj.get("suffix").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `suffix` to be a primitive type in the JSON string but got `%s`", jsonObj.get("suffix").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!Exchange.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'Exchange' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<Exchange> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(Exchange.class));

       return (TypeAdapter<T>) new TypeAdapter<Exchange>() {
           @Override
           public void write(JsonWriter out, Exchange value) throws IOException {
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
           public Exchange read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             Exchange instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of Exchange given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of Exchange
  * @throws IOException if the JSON string is invalid with respect to Exchange
  */
  public static Exchange fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, Exchange.class);
  }

 /**
  * Convert an instance of Exchange to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

