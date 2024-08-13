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
 * Data to login a user via SnapTrade Partner
 */
@ApiModel(description = "Data to login a user via SnapTrade Partner")@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class SnapTradeLoginUserRequestBody {
  public static final String SERIALIZED_NAME_BROKER = "broker";
  @SerializedName(SERIALIZED_NAME_BROKER)
  private String broker;

  public static final String SERIALIZED_NAME_IMMEDIATE_REDIRECT = "immediateRedirect";
  @SerializedName(SERIALIZED_NAME_IMMEDIATE_REDIRECT)
  private Boolean immediateRedirect;

  public static final String SERIALIZED_NAME_CUSTOM_REDIRECT = "customRedirect";
  @SerializedName(SERIALIZED_NAME_CUSTOM_REDIRECT)
  private String customRedirect;

  public static final String SERIALIZED_NAME_RECONNECT = "reconnect";
  @SerializedName(SERIALIZED_NAME_RECONNECT)
  private String reconnect;

  /**
   * Sets whether the connection should be read or trade
   */
  @JsonAdapter(ConnectionTypeEnum.Adapter.class)
 public enum ConnectionTypeEnum {
    READ("read"),
    
    TRADE("trade");

    private String value;

    ConnectionTypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static ConnectionTypeEnum fromValue(String value) {
      for (ConnectionTypeEnum b : ConnectionTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<ConnectionTypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final ConnectionTypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public ConnectionTypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return ConnectionTypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_CONNECTION_TYPE = "connectionType";
  @SerializedName(SERIALIZED_NAME_CONNECTION_TYPE)
  private ConnectionTypeEnum connectionType;

  /**
   * Sets the version of the connection portal to render, with a default to &#39;v3&#39;
   */
  @JsonAdapter(ConnectionPortalVersionEnum.Adapter.class)
 public enum ConnectionPortalVersionEnum {
    V2("v2"),
    
    V3("v3");

    private String value;

    ConnectionPortalVersionEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static ConnectionPortalVersionEnum fromValue(String value) {
      for (ConnectionPortalVersionEnum b : ConnectionPortalVersionEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<ConnectionPortalVersionEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final ConnectionPortalVersionEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public ConnectionPortalVersionEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return ConnectionPortalVersionEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_CONNECTION_PORTAL_VERSION = "connectionPortalVersion";
  @SerializedName(SERIALIZED_NAME_CONNECTION_PORTAL_VERSION)
  private ConnectionPortalVersionEnum connectionPortalVersion;

  public SnapTradeLoginUserRequestBody() {
  }

  public SnapTradeLoginUserRequestBody broker(String broker) {
    
    
    
    
    this.broker = broker;
    return this;
  }

   /**
   * Slug of the brokerage to connect the user to. See [this document](https://snaptrade.notion.site/SnapTrade-Brokerage-Integrations-f83946a714a84c3caf599f6a945f0ead) for a list of supported brokerages and their slugs.
   * @return broker
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "ALPACA", value = "Slug of the brokerage to connect the user to. See [this document](https://snaptrade.notion.site/SnapTrade-Brokerage-Integrations-f83946a714a84c3caf599f6a945f0ead) for a list of supported brokerages and their slugs.")

  public String getBroker() {
    return broker;
  }


  public void setBroker(String broker) {
    
    
    
    this.broker = broker;
  }


  public SnapTradeLoginUserRequestBody immediateRedirect(Boolean immediateRedirect) {
    
    
    
    
    this.immediateRedirect = immediateRedirect;
    return this;
  }

   /**
   * When set to True, user will be redirected back to the partner&#39;s site instead of the connection portal
   * @return immediateRedirect
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "true", value = "When set to True, user will be redirected back to the partner's site instead of the connection portal")

  public Boolean getImmediateRedirect() {
    return immediateRedirect;
  }


  public void setImmediateRedirect(Boolean immediateRedirect) {
    
    
    
    this.immediateRedirect = immediateRedirect;
  }


  public SnapTradeLoginUserRequestBody customRedirect(String customRedirect) {
    
    
    
    
    this.customRedirect = customRedirect;
    return this;
  }

   /**
   * URL to redirect the user to after the user connects their brokerage account
   * @return customRedirect
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "https://snaptrade.com", value = "URL to redirect the user to after the user connects their brokerage account")

  public String getCustomRedirect() {
    return customRedirect;
  }


  public void setCustomRedirect(String customRedirect) {
    
    
    
    this.customRedirect = customRedirect;
  }


  public SnapTradeLoginUserRequestBody reconnect(String reconnect) {
    
    
    
    
    this.reconnect = reconnect;
    return this;
  }

   /**
   * The UUID of the brokerage connection to be reconnected. This parameter should be left empty unless you are reconnecting a disabled connection. See ‘Reconnecting Accounts’ for more information.
   * @return reconnect
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "8b5f262d-4bb9-365d-888a-202bd3b15fa1", value = "The UUID of the brokerage connection to be reconnected. This parameter should be left empty unless you are reconnecting a disabled connection. See ‘Reconnecting Accounts’ for more information.")

  public String getReconnect() {
    return reconnect;
  }


  public void setReconnect(String reconnect) {
    
    
    
    this.reconnect = reconnect;
  }


  public SnapTradeLoginUserRequestBody connectionType(ConnectionTypeEnum connectionType) {
    
    
    
    
    this.connectionType = connectionType;
    return this;
  }

   /**
   * Sets whether the connection should be read or trade
   * @return connectionType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Sets whether the connection should be read or trade")

  public ConnectionTypeEnum getConnectionType() {
    return connectionType;
  }


  public void setConnectionType(ConnectionTypeEnum connectionType) {
    
    
    
    this.connectionType = connectionType;
  }


  public SnapTradeLoginUserRequestBody connectionPortalVersion(ConnectionPortalVersionEnum connectionPortalVersion) {
    
    
    
    
    this.connectionPortalVersion = connectionPortalVersion;
    return this;
  }

   /**
   * Sets the version of the connection portal to render, with a default to &#39;v3&#39;
   * @return connectionPortalVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Sets the version of the connection portal to render, with a default to 'v3'")

  public ConnectionPortalVersionEnum getConnectionPortalVersion() {
    return connectionPortalVersion;
  }


  public void setConnectionPortalVersion(ConnectionPortalVersionEnum connectionPortalVersion) {
    
    
    
    this.connectionPortalVersion = connectionPortalVersion;
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
   * @return the SnapTradeLoginUserRequestBody instance itself
   */
  public SnapTradeLoginUserRequestBody putAdditionalProperty(String key, Object value) {
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
    SnapTradeLoginUserRequestBody snapTradeLoginUserRequestBody = (SnapTradeLoginUserRequestBody) o;
    return Objects.equals(this.broker, snapTradeLoginUserRequestBody.broker) &&
        Objects.equals(this.immediateRedirect, snapTradeLoginUserRequestBody.immediateRedirect) &&
        Objects.equals(this.customRedirect, snapTradeLoginUserRequestBody.customRedirect) &&
        Objects.equals(this.reconnect, snapTradeLoginUserRequestBody.reconnect) &&
        Objects.equals(this.connectionType, snapTradeLoginUserRequestBody.connectionType) &&
        Objects.equals(this.connectionPortalVersion, snapTradeLoginUserRequestBody.connectionPortalVersion)&&
        Objects.equals(this.additionalProperties, snapTradeLoginUserRequestBody.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(broker, immediateRedirect, customRedirect, reconnect, connectionType, connectionPortalVersion, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SnapTradeLoginUserRequestBody {\n");
    sb.append("    broker: ").append(toIndentedString(broker)).append("\n");
    sb.append("    immediateRedirect: ").append(toIndentedString(immediateRedirect)).append("\n");
    sb.append("    customRedirect: ").append(toIndentedString(customRedirect)).append("\n");
    sb.append("    reconnect: ").append(toIndentedString(reconnect)).append("\n");
    sb.append("    connectionType: ").append(toIndentedString(connectionType)).append("\n");
    sb.append("    connectionPortalVersion: ").append(toIndentedString(connectionPortalVersion)).append("\n");
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
    openapiFields.add("broker");
    openapiFields.add("immediateRedirect");
    openapiFields.add("customRedirect");
    openapiFields.add("reconnect");
    openapiFields.add("connectionType");
    openapiFields.add("connectionPortalVersion");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to SnapTradeLoginUserRequestBody
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!SnapTradeLoginUserRequestBody.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in SnapTradeLoginUserRequestBody is not found in the empty JSON string", SnapTradeLoginUserRequestBody.openapiRequiredFields.toString()));
        }
      }
      if ((jsonObj.get("broker") != null && !jsonObj.get("broker").isJsonNull()) && !jsonObj.get("broker").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `broker` to be a primitive type in the JSON string but got `%s`", jsonObj.get("broker").toString()));
      }
      if ((jsonObj.get("customRedirect") != null && !jsonObj.get("customRedirect").isJsonNull()) && !jsonObj.get("customRedirect").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `customRedirect` to be a primitive type in the JSON string but got `%s`", jsonObj.get("customRedirect").toString()));
      }
      if ((jsonObj.get("reconnect") != null && !jsonObj.get("reconnect").isJsonNull()) && !jsonObj.get("reconnect").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `reconnect` to be a primitive type in the JSON string but got `%s`", jsonObj.get("reconnect").toString()));
      }
      if ((jsonObj.get("connectionType") != null && !jsonObj.get("connectionType").isJsonNull()) && !jsonObj.get("connectionType").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `connectionType` to be a primitive type in the JSON string but got `%s`", jsonObj.get("connectionType").toString()));
      }
      if ((jsonObj.get("connectionPortalVersion") != null && !jsonObj.get("connectionPortalVersion").isJsonNull()) && !jsonObj.get("connectionPortalVersion").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `connectionPortalVersion` to be a primitive type in the JSON string but got `%s`", jsonObj.get("connectionPortalVersion").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!SnapTradeLoginUserRequestBody.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'SnapTradeLoginUserRequestBody' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<SnapTradeLoginUserRequestBody> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(SnapTradeLoginUserRequestBody.class));

       return (TypeAdapter<T>) new TypeAdapter<SnapTradeLoginUserRequestBody>() {
           @Override
           public void write(JsonWriter out, SnapTradeLoginUserRequestBody value) throws IOException {
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
                 else if (entry.getValue() == null) {
                   obj.addProperty(entry.getKey(), (String) null);
                 } else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public SnapTradeLoginUserRequestBody read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             SnapTradeLoginUserRequestBody instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of SnapTradeLoginUserRequestBody given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of SnapTradeLoginUserRequestBody
  * @throws IOException if the JSON string is invalid with respect to SnapTradeLoginUserRequestBody
  */
  public static SnapTradeLoginUserRequestBody fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, SnapTradeLoginUserRequestBody.class);
  }

 /**
  * Convert an instance of SnapTradeLoginUserRequestBody to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

