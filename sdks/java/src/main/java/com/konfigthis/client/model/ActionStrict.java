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
import io.swagger.annotations.ApiModel;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * The action describes the intent or side of a trade. This is either &#x60;BUY&#x60; or &#x60;SELL&#x60;.
 */
@JsonAdapter(ActionStrict.Adapter.class)public enum ActionStrict {
  
  BUY("BUY"),
  
  SELL("SELL");

  private String value;

  ActionStrict(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static ActionStrict fromValue(String value) {
    for (ActionStrict b : ActionStrict.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<ActionStrict> {
    @Override
    public void write(final JsonWriter jsonWriter, final ActionStrict enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public ActionStrict read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return ActionStrict.fromValue(value);
    }
  }
}

