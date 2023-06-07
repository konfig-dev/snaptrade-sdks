/*
 * SnapTrade
 *
 * Connect brokerage accounts to your app for live positions and trading
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: api@snaptrade.com
 * Generated by: https://konfigthis.com
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = SnapTrade.Net.Client.OpenAPIDateConverter;

namespace SnapTrade.Net.Model
{
    /// <summary>
    /// OptionsPlaceOptionStrategyRequest
    /// </summary>
    [DataContract(Name = "Options_placeOptionStrategy_request")]
    public partial class OptionsPlaceOptionStrategyRequest : IEquatable<OptionsPlaceOptionStrategyRequest>, IValidatableObject
    {
        /// <summary>
        /// Defines OrderType
        /// </summary>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum OrderTypeEnum
        {
            /// <summary>
            /// Enum Limit for value: Limit
            /// </summary>
            [EnumMember(Value = "Limit")]
            Limit = 1,

            /// <summary>
            /// Enum Market for value: Market
            /// </summary>
            [EnumMember(Value = "Market")]
            Market = 2,

            /// <summary>
            /// Enum NetDebit for value: NetDebit
            /// </summary>
            [EnumMember(Value = "NetDebit")]
            NetDebit = 3,

            /// <summary>
            /// Enum NetCredit for value: NetCredit
            /// </summary>
            [EnumMember(Value = "NetCredit")]
            NetCredit = 4

        }


        /// <summary>
        /// Gets or Sets OrderType
        /// </summary>
        [DataMember(Name = "order_type", IsRequired = true, EmitDefaultValue = true)]
        public OrderTypeEnum OrderType { get; set; }
        /// <summary>
        /// Defines TimeInForce
        /// </summary>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum TimeInForceEnum
        {
            /// <summary>
            /// Enum DAY for value: DAY
            /// </summary>
            [EnumMember(Value = "DAY")]
            DAY = 1,

            /// <summary>
            /// Enum GTC for value: GTC
            /// </summary>
            [EnumMember(Value = "GTC")]
            GTC = 2

        }


        /// <summary>
        /// Gets or Sets TimeInForce
        /// </summary>
        [DataMember(Name = "time_in_force", IsRequired = true, EmitDefaultValue = true)]
        public TimeInForceEnum TimeInForce { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="OptionsPlaceOptionStrategyRequest" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected OptionsPlaceOptionStrategyRequest() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="OptionsPlaceOptionStrategyRequest" /> class.
        /// </summary>
        /// <param name="orderType">orderType (required).</param>
        /// <param name="timeInForce">timeInForce (required).</param>
        /// <param name="price">Trade Price if limit or stop limit order (required).</param>
        public OptionsPlaceOptionStrategyRequest(OrderTypeEnum orderType = default(OrderTypeEnum), TimeInForceEnum timeInForce = default(TimeInForceEnum), decimal? price = default(decimal?))
        {
            this.OrderType = orderType;
            this.TimeInForce = timeInForce;
            // to ensure "price" is required (not null)
            if (price == null)
            {
                throw new ArgumentNullException("price is a required property for OptionsPlaceOptionStrategyRequest and cannot be null");
            }
            this.Price = price;
        }

        /// <summary>
        /// Trade Price if limit or stop limit order
        /// </summary>
        /// <value>Trade Price if limit or stop limit order</value>
        [DataMember(Name = "price", IsRequired = true, EmitDefaultValue = true)]
        public decimal? Price { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class OptionsPlaceOptionStrategyRequest {\n");
            sb.Append("  OrderType: ").Append(OrderType).Append("\n");
            sb.Append("  TimeInForce: ").Append(TimeInForce).Append("\n");
            sb.Append("  Price: ").Append(Price).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as OptionsPlaceOptionStrategyRequest);
        }

        /// <summary>
        /// Returns true if OptionsPlaceOptionStrategyRequest instances are equal
        /// </summary>
        /// <param name="input">Instance of OptionsPlaceOptionStrategyRequest to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OptionsPlaceOptionStrategyRequest input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.OrderType == input.OrderType ||
                    this.OrderType.Equals(input.OrderType)
                ) && 
                (
                    this.TimeInForce == input.TimeInForce ||
                    this.TimeInForce.Equals(input.TimeInForce)
                ) && 
                (
                    this.Price == input.Price ||
                    (this.Price != null &&
                    this.Price.Equals(input.Price))
                );
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                hashCode = (hashCode * 59) + this.OrderType.GetHashCode();
                hashCode = (hashCode * 59) + this.TimeInForce.GetHashCode();
                if (this.Price != null)
                {
                    hashCode = (hashCode * 59) + this.Price.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
