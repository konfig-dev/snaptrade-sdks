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
    /// Option Holdings
    /// </summary>
    [DataContract(Name = "OptionsHoldings")]
    public partial class OptionsHoldings : IEquatable<OptionsHoldings>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="OptionsHoldings" /> class.
        /// </summary>
        /// <param name="id">Options information.</param>
        /// <param name="symbol">symbol.</param>
        /// <param name="optionSymbol">optionSymbol.</param>
        /// <param name="price">Trade Price if limit or stop limit order.</param>
        /// <param name="currency">currency.</param>
        /// <param name="averagePurchasePrice">Average purchase price for this position.</param>
        public OptionsHoldings(string id = default(string), Guid symbol = default(Guid), OptionsSymbol optionSymbol = default(OptionsSymbol), decimal price = default(decimal), Currency currency = default(Currency), decimal averagePurchasePrice = default(decimal)) : base()
        {
            this.Id = id;
            this.Symbol = symbol;
            this.OptionSymbol = optionSymbol;
            this.Price = price;
            this.Currency = currency;
            this.AveragePurchasePrice = averagePurchasePrice;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Options information
        /// </summary>
        /// <value>Options information</value>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// Gets or Sets Symbol
        /// </summary>
        [DataMember(Name = "symbol", EmitDefaultValue = false)]
        public Guid Symbol { get; set; }

        /// <summary>
        /// Gets or Sets OptionSymbol
        /// </summary>
        [DataMember(Name = "option_symbol", EmitDefaultValue = false)]
        public OptionsSymbol OptionSymbol { get; set; }

        /// <summary>
        /// Trade Price if limit or stop limit order
        /// </summary>
        /// <value>Trade Price if limit or stop limit order</value>
        [DataMember(Name = "price", EmitDefaultValue = false)]
        public decimal Price { get; set; }

        /// <summary>
        /// Gets or Sets Currency
        /// </summary>
        [DataMember(Name = "currency", EmitDefaultValue = false)]
        public Currency Currency { get; set; }

        /// <summary>
        /// Average purchase price for this position
        /// </summary>
        /// <value>Average purchase price for this position</value>
        [DataMember(Name = "average_purchase_price", EmitDefaultValue = false)]
        public decimal AveragePurchasePrice { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class OptionsHoldings {\n");
            sb.Append("  ").Append(base.ToString().Replace("\n", "\n  ")).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Symbol: ").Append(Symbol).Append("\n");
            sb.Append("  OptionSymbol: ").Append(OptionSymbol).Append("\n");
            sb.Append("  Price: ").Append(Price).Append("\n");
            sb.Append("  Currency: ").Append(Currency).Append("\n");
            sb.Append("  AveragePurchasePrice: ").Append(AveragePurchasePrice).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public string ToJson()
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
            return this.Equals(input as OptionsHoldings);
        }

        /// <summary>
        /// Returns true if OptionsHoldings instances are equal
        /// </summary>
        /// <param name="input">Instance of OptionsHoldings to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OptionsHoldings input)
        {
            if (input == null)
            {
                return false;
            }
            return base.Equals(input) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && base.Equals(input) && 
                (
                    this.Symbol == input.Symbol ||
                    (this.Symbol != null &&
                    this.Symbol.Equals(input.Symbol))
                ) && base.Equals(input) && 
                (
                    this.OptionSymbol == input.OptionSymbol ||
                    (this.OptionSymbol != null &&
                    this.OptionSymbol.Equals(input.OptionSymbol))
                ) && base.Equals(input) && 
                (
                    this.Price == input.Price ||
                    this.Price.Equals(input.Price)
                ) && base.Equals(input) && 
                (
                    this.Currency == input.Currency ||
                    (this.Currency != null &&
                    this.Currency.Equals(input.Currency))
                ) && base.Equals(input) && 
                (
                    this.AveragePurchasePrice == input.AveragePurchasePrice ||
                    this.AveragePurchasePrice.Equals(input.AveragePurchasePrice)
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = base.GetHashCode();
                if (this.Id != null)
                {
                    hashCode = (hashCode * 59) + this.Id.GetHashCode();
                }
                if (this.Symbol != null)
                {
                    hashCode = (hashCode * 59) + this.Symbol.GetHashCode();
                }
                if (this.OptionSymbol != null)
                {
                    hashCode = (hashCode * 59) + this.OptionSymbol.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Price.GetHashCode();
                if (this.Currency != null)
                {
                    hashCode = (hashCode * 59) + this.Currency.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.AveragePurchasePrice.GetHashCode();
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
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
