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
    /// StrategyImpactLegsInner
    /// </summary>
    [DataContract(Name = "StrategyImpact_legs_inner")]
    public partial class StrategyImpactLegsInner : IEquatable<StrategyImpactLegsInner>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="StrategyImpactLegsInner" /> class.
        /// </summary>
        /// <param name="legId">legId.</param>
        /// <param name="symbol">symbol.</param>
        /// <param name="symbolId">symbolId.</param>
        /// <param name="legRatioQuantity">legRatioQuantity.</param>
        /// <param name="side">side.</param>
        /// <param name="avgExecPrice">avgExecPrice.</param>
        /// <param name="lastExecPrice">lastExecPrice.</param>
        public StrategyImpactLegsInner(int legId = default(int), string symbol = default(string), int symbolId = default(int), int legRatioQuantity = default(int), string side = default(string), string avgExecPrice = default(string), string lastExecPrice = default(string)) : base()
        {
            this.LegId = legId;
            this.Symbol = symbol;
            this.SymbolId = symbolId;
            this.LegRatioQuantity = legRatioQuantity;
            this.Side = side;
            this.AvgExecPrice = avgExecPrice;
            this.LastExecPrice = lastExecPrice;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets LegId
        /// </summary>
        [DataMember(Name = "legId", EmitDefaultValue = false)]
        public int LegId { get; set; }

        /// <summary>
        /// Gets or Sets Symbol
        /// </summary>
        [DataMember(Name = "symbol", EmitDefaultValue = false)]
        public string Symbol { get; set; }

        /// <summary>
        /// Gets or Sets SymbolId
        /// </summary>
        [DataMember(Name = "symbolId", EmitDefaultValue = false)]
        public int SymbolId { get; set; }

        /// <summary>
        /// Gets or Sets LegRatioQuantity
        /// </summary>
        [DataMember(Name = "legRatioQuantity", EmitDefaultValue = false)]
        public int LegRatioQuantity { get; set; }

        /// <summary>
        /// Gets or Sets Side
        /// </summary>
        [DataMember(Name = "side", EmitDefaultValue = false)]
        public string Side { get; set; }

        /// <summary>
        /// Gets or Sets AvgExecPrice
        /// </summary>
        [DataMember(Name = "avgExecPrice", EmitDefaultValue = false)]
        public string AvgExecPrice { get; set; }

        /// <summary>
        /// Gets or Sets LastExecPrice
        /// </summary>
        [DataMember(Name = "lastExecPrice", EmitDefaultValue = false)]
        public string LastExecPrice { get; set; }

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
            sb.Append("class StrategyImpactLegsInner {\n");
            sb.Append("  ").Append(base.ToString().Replace("\n", "\n  ")).Append("\n");
            sb.Append("  LegId: ").Append(LegId).Append("\n");
            sb.Append("  Symbol: ").Append(Symbol).Append("\n");
            sb.Append("  SymbolId: ").Append(SymbolId).Append("\n");
            sb.Append("  LegRatioQuantity: ").Append(LegRatioQuantity).Append("\n");
            sb.Append("  Side: ").Append(Side).Append("\n");
            sb.Append("  AvgExecPrice: ").Append(AvgExecPrice).Append("\n");
            sb.Append("  LastExecPrice: ").Append(LastExecPrice).Append("\n");
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
            return this.Equals(input as StrategyImpactLegsInner);
        }

        /// <summary>
        /// Returns true if StrategyImpactLegsInner instances are equal
        /// </summary>
        /// <param name="input">Instance of StrategyImpactLegsInner to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(StrategyImpactLegsInner input)
        {
            if (input == null)
            {
                return false;
            }
            return base.Equals(input) && 
                (
                    this.LegId == input.LegId ||
                    this.LegId.Equals(input.LegId)
                ) && base.Equals(input) && 
                (
                    this.Symbol == input.Symbol ||
                    (this.Symbol != null &&
                    this.Symbol.Equals(input.Symbol))
                ) && base.Equals(input) && 
                (
                    this.SymbolId == input.SymbolId ||
                    this.SymbolId.Equals(input.SymbolId)
                ) && base.Equals(input) && 
                (
                    this.LegRatioQuantity == input.LegRatioQuantity ||
                    this.LegRatioQuantity.Equals(input.LegRatioQuantity)
                ) && base.Equals(input) && 
                (
                    this.Side == input.Side ||
                    (this.Side != null &&
                    this.Side.Equals(input.Side))
                ) && base.Equals(input) && 
                (
                    this.AvgExecPrice == input.AvgExecPrice ||
                    (this.AvgExecPrice != null &&
                    this.AvgExecPrice.Equals(input.AvgExecPrice))
                ) && base.Equals(input) && 
                (
                    this.LastExecPrice == input.LastExecPrice ||
                    (this.LastExecPrice != null &&
                    this.LastExecPrice.Equals(input.LastExecPrice))
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
                hashCode = (hashCode * 59) + this.LegId.GetHashCode();
                if (this.Symbol != null)
                {
                    hashCode = (hashCode * 59) + this.Symbol.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.SymbolId.GetHashCode();
                hashCode = (hashCode * 59) + this.LegRatioQuantity.GetHashCode();
                if (this.Side != null)
                {
                    hashCode = (hashCode * 59) + this.Side.GetHashCode();
                }
                if (this.AvgExecPrice != null)
                {
                    hashCode = (hashCode * 59) + this.AvgExecPrice.GetHashCode();
                }
                if (this.LastExecPrice != null)
                {
                    hashCode = (hashCode * 59) + this.LastExecPrice.GetHashCode();
                }
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
