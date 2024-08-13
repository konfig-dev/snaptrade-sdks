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
import com.konfigthis.client.model.Balance;
import com.konfigthis.client.model.Position;
import com.konfigthis.client.model.SnapTradeHoldingsAccount;
import com.konfigthis.client.model.SnapTradeHoldingsTotalValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
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
 * Account Holdings
 */
@ApiModel(description = "Account Holdings")@javax.annotation.Generated(value = "Generated by https://konfigthis.com")
public class AccountHoldings {
  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private SnapTradeHoldingsAccount account;

  public static final String SERIALIZED_NAME_BALANCES = "balances";
  @SerializedName(SERIALIZED_NAME_BALANCES)
  private List<Balance> balances = null;

  public static final String SERIALIZED_NAME_POSITIONS = "positions";
  @SerializedName(SERIALIZED_NAME_POSITIONS)
  private List<Position> positions = null;

  public static final String SERIALIZED_NAME_TOTAL_VALUE = "total_value";
  @SerializedName(SERIALIZED_NAME_TOTAL_VALUE)
  private SnapTradeHoldingsTotalValue totalValue;

  public AccountHoldings() {
  }

  public AccountHoldings account(SnapTradeHoldingsAccount account) {
    
    
    
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public SnapTradeHoldingsAccount getAccount() {
    return account;
  }


  public void setAccount(SnapTradeHoldingsAccount account) {
    
    
    
    this.account = account;
  }


  public AccountHoldings balances(List<Balance> balances) {
    
    
    
    
    this.balances = balances;
    return this;
  }

  public AccountHoldings addBalancesItem(Balance balancesItem) {
    if (this.balances == null) {
      this.balances = new ArrayList<>();
    }
    this.balances.add(balancesItem);
    return this;
  }

   /**
   * Get balances
   * @return balances
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Balance> getBalances() {
    return balances;
  }


  public void setBalances(List<Balance> balances) {
    
    
    
    this.balances = balances;
  }


  public AccountHoldings positions(List<Position> positions) {
    
    
    
    
    this.positions = positions;
    return this;
  }

  public AccountHoldings addPositionsItem(Position positionsItem) {
    if (this.positions == null) {
      this.positions = new ArrayList<>();
    }
    this.positions.add(positionsItem);
    return this;
  }

   /**
   * Get positions
   * @return positions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Position> getPositions() {
    return positions;
  }


  public void setPositions(List<Position> positions) {
    
    
    
    this.positions = positions;
  }


  public AccountHoldings totalValue(SnapTradeHoldingsTotalValue totalValue) {
    
    
    
    
    this.totalValue = totalValue;
    return this;
  }

   /**
   * Get totalValue
   * @return totalValue
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public SnapTradeHoldingsTotalValue getTotalValue() {
    return totalValue;
  }


  public void setTotalValue(SnapTradeHoldingsTotalValue totalValue) {
    
    
    
    this.totalValue = totalValue;
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
   * @return the AccountHoldings instance itself
   */
  public AccountHoldings putAdditionalProperty(String key, Object value) {
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
    AccountHoldings accountHoldings = (AccountHoldings) o;
    return Objects.equals(this.account, accountHoldings.account) &&
        Objects.equals(this.balances, accountHoldings.balances) &&
        Objects.equals(this.positions, accountHoldings.positions) &&
        Objects.equals(this.totalValue, accountHoldings.totalValue)&&
        Objects.equals(this.additionalProperties, accountHoldings.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(account, balances, positions, totalValue, additionalProperties);
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
    sb.append("class AccountHoldings {\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    balances: ").append(toIndentedString(balances)).append("\n");
    sb.append("    positions: ").append(toIndentedString(positions)).append("\n");
    sb.append("    totalValue: ").append(toIndentedString(totalValue)).append("\n");
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
    openapiFields.add("account");
    openapiFields.add("balances");
    openapiFields.add("positions");
    openapiFields.add("total_value");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to AccountHoldings
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!AccountHoldings.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AccountHoldings is not found in the empty JSON string", AccountHoldings.openapiRequiredFields.toString()));
        }
      }
      // validate the optional field `account`
      if (jsonObj.get("account") != null && !jsonObj.get("account").isJsonNull()) {
        SnapTradeHoldingsAccount.validateJsonObject(jsonObj.getAsJsonObject("account"));
      }
      if (jsonObj.get("balances") != null && !jsonObj.get("balances").isJsonNull()) {
        JsonArray jsonArraybalances = jsonObj.getAsJsonArray("balances");
        if (jsonArraybalances != null) {
          // ensure the json data is an array
          if (!jsonObj.get("balances").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `balances` to be an array in the JSON string but got `%s`", jsonObj.get("balances").toString()));
          }

          // validate the optional field `balances` (array)
          for (int i = 0; i < jsonArraybalances.size(); i++) {
            Balance.validateJsonObject(jsonArraybalances.get(i).getAsJsonObject());
          };
        }
      }
      if (jsonObj.get("positions") != null && !jsonObj.get("positions").isJsonNull()) {
        JsonArray jsonArraypositions = jsonObj.getAsJsonArray("positions");
        if (jsonArraypositions != null) {
          // ensure the json data is an array
          if (!jsonObj.get("positions").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `positions` to be an array in the JSON string but got `%s`", jsonObj.get("positions").toString()));
          }

          // validate the optional field `positions` (array)
          for (int i = 0; i < jsonArraypositions.size(); i++) {
            Position.validateJsonObject(jsonArraypositions.get(i).getAsJsonObject());
          };
        }
      }
      // validate the optional field `total_value`
      if (jsonObj.get("total_value") != null && !jsonObj.get("total_value").isJsonNull()) {
        SnapTradeHoldingsTotalValue.validateJsonObject(jsonObj.getAsJsonObject("total_value"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AccountHoldings.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AccountHoldings' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AccountHoldings> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AccountHoldings.class));

       return (TypeAdapter<T>) new TypeAdapter<AccountHoldings>() {
           @Override
           public void write(JsonWriter out, AccountHoldings value) throws IOException {
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
           public AccountHoldings read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             AccountHoldings instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of AccountHoldings given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AccountHoldings
  * @throws IOException if the JSON string is invalid with respect to AccountHoldings
  */
  public static AccountHoldings fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AccountHoldings.class);
  }

 /**
  * Convert an instance of AccountHoldings to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

