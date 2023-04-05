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
    /// Options Symbol
    /// </summary>
    [DataContract(Name = "OptionsSymbol")]
    public partial class OptionsSymbol : IEquatable<OptionsSymbol>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="OptionsSymbol" /> class.
        /// </summary>
        /// <param name="id">id.</param>
        /// <param name="ticker">ticker.</param>
        /// <param name="strikePrice">strikePrice.</param>
        /// <param name="expirationDate">expirationDate.</param>
        /// <param name="isMiniOption">isMiniOption.</param>
        /// <param name="underlyingSymbol">underlyingSymbol.</param>
        /// <param name="localId">localId.</param>
        /// <param name="securityType">securityType.</param>
        /// <param name="listingExchange">listingExchange.</param>
        /// <param name="isQuotable">isQuotable.</param>
        /// <param name="isTradable">isTradable.</param>
        public OptionsSymbol(Guid id = default(Guid), string ticker = default(string), int strikePrice = default(int), string expirationDate = default(string), bool isMiniOption = default(bool), UnderlyingSymbol underlyingSymbol = default(UnderlyingSymbol), string localId = default(string), Object securityType = default(Object), Object listingExchange = default(Object), bool isQuotable = default(bool), bool isTradable = default(bool)) : base()
        {
            this.Id = id;
            this.Ticker = ticker;
            this.StrikePrice = strikePrice;
            this.ExpirationDate = expirationDate;
            this.IsMiniOption = isMiniOption;
            this.UnderlyingSymbol = underlyingSymbol;
            this.LocalId = localId;
            this.SecurityType = securityType;
            this.ListingExchange = listingExchange;
            this.IsQuotable = isQuotable;
            this.IsTradable = isTradable;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public Guid Id { get; set; }

        /// <summary>
        /// Gets or Sets Ticker
        /// </summary>
        [DataMember(Name = "ticker", EmitDefaultValue = false)]
        public string Ticker { get; set; }

        /// <summary>
        /// Gets or Sets StrikePrice
        /// </summary>
        [DataMember(Name = "strike_price", EmitDefaultValue = false)]
        public int StrikePrice { get; set; }

        /// <summary>
        /// Gets or Sets ExpirationDate
        /// </summary>
        [DataMember(Name = "expiration_date", EmitDefaultValue = false)]
        public string ExpirationDate { get; set; }

        /// <summary>
        /// Gets or Sets IsMiniOption
        /// </summary>
        [DataMember(Name = "is_mini_option", EmitDefaultValue = true)]
        public bool IsMiniOption { get; set; }

        /// <summary>
        /// Gets or Sets UnderlyingSymbol
        /// </summary>
        [DataMember(Name = "underlying_symbol", EmitDefaultValue = false)]
        public UnderlyingSymbol UnderlyingSymbol { get; set; }

        /// <summary>
        /// Gets or Sets LocalId
        /// </summary>
        [DataMember(Name = "local_id", EmitDefaultValue = false)]
        public string LocalId { get; set; }

        /// <summary>
        /// Gets or Sets SecurityType
        /// </summary>
        [DataMember(Name = "security_type", EmitDefaultValue = true)]
        public Object SecurityType { get; set; }

        /// <summary>
        /// Gets or Sets ListingExchange
        /// </summary>
        [DataMember(Name = "listing_exchange", EmitDefaultValue = true)]
        public Object ListingExchange { get; set; }

        /// <summary>
        /// Gets or Sets IsQuotable
        /// </summary>
        [DataMember(Name = "is_quotable", EmitDefaultValue = true)]
        public bool IsQuotable { get; set; }

        /// <summary>
        /// Gets or Sets IsTradable
        /// </summary>
        [DataMember(Name = "is_tradable", EmitDefaultValue = true)]
        public bool IsTradable { get; set; }

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
            sb.Append("class OptionsSymbol {\n");
            sb.Append("  ").Append(base.ToString().Replace("\n", "\n  ")).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Ticker: ").Append(Ticker).Append("\n");
            sb.Append("  StrikePrice: ").Append(StrikePrice).Append("\n");
            sb.Append("  ExpirationDate: ").Append(ExpirationDate).Append("\n");
            sb.Append("  IsMiniOption: ").Append(IsMiniOption).Append("\n");
            sb.Append("  UnderlyingSymbol: ").Append(UnderlyingSymbol).Append("\n");
            sb.Append("  LocalId: ").Append(LocalId).Append("\n");
            sb.Append("  SecurityType: ").Append(SecurityType).Append("\n");
            sb.Append("  ListingExchange: ").Append(ListingExchange).Append("\n");
            sb.Append("  IsQuotable: ").Append(IsQuotable).Append("\n");
            sb.Append("  IsTradable: ").Append(IsTradable).Append("\n");
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
            return this.Equals(input as OptionsSymbol);
        }

        /// <summary>
        /// Returns true if OptionsSymbol instances are equal
        /// </summary>
        /// <param name="input">Instance of OptionsSymbol to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OptionsSymbol input)
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
                    this.Ticker == input.Ticker ||
                    (this.Ticker != null &&
                    this.Ticker.Equals(input.Ticker))
                ) && base.Equals(input) && 
                (
                    this.StrikePrice == input.StrikePrice ||
                    this.StrikePrice.Equals(input.StrikePrice)
                ) && base.Equals(input) && 
                (
                    this.ExpirationDate == input.ExpirationDate ||
                    (this.ExpirationDate != null &&
                    this.ExpirationDate.Equals(input.ExpirationDate))
                ) && base.Equals(input) && 
                (
                    this.IsMiniOption == input.IsMiniOption ||
                    this.IsMiniOption.Equals(input.IsMiniOption)
                ) && base.Equals(input) && 
                (
                    this.UnderlyingSymbol == input.UnderlyingSymbol ||
                    (this.UnderlyingSymbol != null &&
                    this.UnderlyingSymbol.Equals(input.UnderlyingSymbol))
                ) && base.Equals(input) && 
                (
                    this.LocalId == input.LocalId ||
                    (this.LocalId != null &&
                    this.LocalId.Equals(input.LocalId))
                ) && base.Equals(input) && 
                (
                    this.SecurityType == input.SecurityType ||
                    (this.SecurityType != null &&
                    this.SecurityType.Equals(input.SecurityType))
                ) && base.Equals(input) && 
                (
                    this.ListingExchange == input.ListingExchange ||
                    (this.ListingExchange != null &&
                    this.ListingExchange.Equals(input.ListingExchange))
                ) && base.Equals(input) && 
                (
                    this.IsQuotable == input.IsQuotable ||
                    this.IsQuotable.Equals(input.IsQuotable)
                ) && base.Equals(input) && 
                (
                    this.IsTradable == input.IsTradable ||
                    this.IsTradable.Equals(input.IsTradable)
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
                if (this.Ticker != null)
                {
                    hashCode = (hashCode * 59) + this.Ticker.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.StrikePrice.GetHashCode();
                if (this.ExpirationDate != null)
                {
                    hashCode = (hashCode * 59) + this.ExpirationDate.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.IsMiniOption.GetHashCode();
                if (this.UnderlyingSymbol != null)
                {
                    hashCode = (hashCode * 59) + this.UnderlyingSymbol.GetHashCode();
                }
                if (this.LocalId != null)
                {
                    hashCode = (hashCode * 59) + this.LocalId.GetHashCode();
                }
                if (this.SecurityType != null)
                {
                    hashCode = (hashCode * 59) + this.SecurityType.GetHashCode();
                }
                if (this.ListingExchange != null)
                {
                    hashCode = (hashCode * 59) + this.ListingExchange.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.IsQuotable.GetHashCode();
                hashCode = (hashCode * 59) + this.IsTradable.GetHashCode();
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
