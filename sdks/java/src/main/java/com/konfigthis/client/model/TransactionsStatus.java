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
 * Status of account transaction sync
 */
@ApiModel(description = "Status of account transaction sync")@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class TransactionsStatus {
  public static final String SERIALIZED_NAME_INITIAL_SYNC_COMPLETED = "initial_sync_completed";
  @SerializedName(SERIALIZED_NAME_INITIAL_SYNC_COMPLETED)
  private Boolean initialSyncCompleted;

  public static final String SERIALIZED_NAME_LAST_SUCCESSFUL_SYNC = "last_successful_sync";
  @SerializedName(SERIALIZED_NAME_LAST_SUCCESSFUL_SYNC)
  private LocalDate lastSuccessfulSync;

  public static final String SERIALIZED_NAME_FIRST_TRANSACTION_DATE = "first_transaction_date";
  @SerializedName(SERIALIZED_NAME_FIRST_TRANSACTION_DATE)
  private LocalDate firstTransactionDate;

  public TransactionsStatus() {
  }

  public TransactionsStatus initialSyncCompleted(Boolean initialSyncCompleted) {
    
    
    
    
    this.initialSyncCompleted = initialSyncCompleted;
    return this;
  }

   /**
   * Get initialSyncCompleted
   * @return initialSyncCompleted
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getInitialSyncCompleted() {
    return initialSyncCompleted;
  }


  public void setInitialSyncCompleted(Boolean initialSyncCompleted) {
    
    
    
    this.initialSyncCompleted = initialSyncCompleted;
  }


  public TransactionsStatus lastSuccessfulSync(LocalDate lastSuccessfulSync) {
    
    
    
    
    this.lastSuccessfulSync = lastSuccessfulSync;
    return this;
  }

   /**
   * Date in YYYY-MM-DD format or null
   * @return lastSuccessfulSync
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Sun Jan 23 16:00:00 PST 2022", value = "Date in YYYY-MM-DD format or null")

  public LocalDate getLastSuccessfulSync() {
    return lastSuccessfulSync;
  }


  public void setLastSuccessfulSync(LocalDate lastSuccessfulSync) {
    
    
    
    this.lastSuccessfulSync = lastSuccessfulSync;
  }


  public TransactionsStatus firstTransactionDate(LocalDate firstTransactionDate) {
    
    
    
    
    this.firstTransactionDate = firstTransactionDate;
    return this;
  }

   /**
   * Date in YYYY-MM-DD format or null
   * @return firstTransactionDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Sun Jan 23 16:00:00 PST 2022", value = "Date in YYYY-MM-DD format or null")

  public LocalDate getFirstTransactionDate() {
    return firstTransactionDate;
  }


  public void setFirstTransactionDate(LocalDate firstTransactionDate) {
    
    
    
    this.firstTransactionDate = firstTransactionDate;
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
   * @return the TransactionsStatus instance itself
   */
  public TransactionsStatus putAdditionalProperty(String key, Object value) {
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
    TransactionsStatus transactionsStatus = (TransactionsStatus) o;
    return Objects.equals(this.initialSyncCompleted, transactionsStatus.initialSyncCompleted) &&
        Objects.equals(this.lastSuccessfulSync, transactionsStatus.lastSuccessfulSync) &&
        Objects.equals(this.firstTransactionDate, transactionsStatus.firstTransactionDate)&&
        Objects.equals(this.additionalProperties, transactionsStatus.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(initialSyncCompleted, lastSuccessfulSync, firstTransactionDate, additionalProperties);
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
    sb.append("class TransactionsStatus {\n");
    sb.append("    initialSyncCompleted: ").append(toIndentedString(initialSyncCompleted)).append("\n");
    sb.append("    lastSuccessfulSync: ").append(toIndentedString(lastSuccessfulSync)).append("\n");
    sb.append("    firstTransactionDate: ").append(toIndentedString(firstTransactionDate)).append("\n");
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
    openapiFields.add("initial_sync_completed");
    openapiFields.add("last_successful_sync");
    openapiFields.add("first_transaction_date");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to TransactionsStatus
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!TransactionsStatus.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in TransactionsStatus is not found in the empty JSON string", TransactionsStatus.openapiRequiredFields.toString()));
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!TransactionsStatus.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'TransactionsStatus' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<TransactionsStatus> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(TransactionsStatus.class));

       return (TypeAdapter<T>) new TypeAdapter<TransactionsStatus>() {
           @Override
           public void write(JsonWriter out, TransactionsStatus value) throws IOException {
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
           public TransactionsStatus read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             TransactionsStatus instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of TransactionsStatus given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of TransactionsStatus
  * @throws IOException if the JSON string is invalid with respect to TransactionsStatus
  */
  public static TransactionsStatus fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, TransactionsStatus.class);
  }

 /**
  * Convert an instance of TransactionsStatus to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

