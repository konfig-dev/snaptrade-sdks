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
    /// A transaction or activity from an institution
    /// </summary>
    [DataContract(Name = "UniversalActivity")]
    public partial class UniversalActivity : IEquatable<UniversalActivity>, IValidatableObject
    {
        /// <summary>
        /// Defines Type
        /// </summary>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum TypeEnum
        {
            /// <summary>
            /// Enum DIVIDEND for value: DIVIDEND
            /// </summary>
            [EnumMember(Value = "DIVIDEND")]
            DIVIDEND = 1,

            /// <summary>
            /// Enum BUY for value: BUY
            /// </summary>
            [EnumMember(Value = "BUY")]
            BUY = 2,

            /// <summary>
            /// Enum SELL for value: SELL
            /// </summary>
            [EnumMember(Value = "SELL")]
            SELL = 3,

            /// <summary>
            /// Enum CONTRIBUTION for value: CONTRIBUTION
            /// </summary>
            [EnumMember(Value = "CONTRIBUTION")]
            CONTRIBUTION = 4,

            /// <summary>
            /// Enum WITHDRAWAL for value: WITHDRAWAL
            /// </summary>
            [EnumMember(Value = "WITHDRAWAL")]
            WITHDRAWAL = 5,

            /// <summary>
            /// Enum EXTERNALASSETTRANSFERIN for value: EXTERNAL_ASSET_TRANSFER_IN
            /// </summary>
            [EnumMember(Value = "EXTERNAL_ASSET_TRANSFER_IN")]
            EXTERNALASSETTRANSFERIN = 6,

            /// <summary>
            /// Enum EXTERNALASSETTRANSFEROUT for value: EXTERNAL_ASSET_TRANSFER_OUT
            /// </summary>
            [EnumMember(Value = "EXTERNAL_ASSET_TRANSFER_OUT")]
            EXTERNALASSETTRANSFEROUT = 7,

            /// <summary>
            /// Enum INTERNALCASHTRANSFERIN for value: INTERNAL_CASH_TRANSFER_IN
            /// </summary>
            [EnumMember(Value = "INTERNAL_CASH_TRANSFER_IN")]
            INTERNALCASHTRANSFERIN = 8,

            /// <summary>
            /// Enum INTERNALCASHTRANSFEROUT for value: INTERNAL_CASH_TRANSFER_OUT
            /// </summary>
            [EnumMember(Value = "INTERNAL_CASH_TRANSFER_OUT")]
            INTERNALCASHTRANSFEROUT = 9,

            /// <summary>
            /// Enum INTERNALASSETTRANSFERIN for value: INTERNAL_ASSET_TRANSFER_IN
            /// </summary>
            [EnumMember(Value = "INTERNAL_ASSET_TRANSFER_IN")]
            INTERNALASSETTRANSFERIN = 10,

            /// <summary>
            /// Enum INTERNALASSETTRANSFEROUT for value: INTERNAL_ASSET_TRANSFER_OUT
            /// </summary>
            [EnumMember(Value = "INTERNAL_ASSET_TRANSFER_OUT")]
            INTERNALASSETTRANSFEROUT = 11,

            /// <summary>
            /// Enum INTEREST for value: INTEREST
            /// </summary>
            [EnumMember(Value = "INTEREST")]
            INTEREST = 12,

            /// <summary>
            /// Enum REBATE for value: REBATE
            /// </summary>
            [EnumMember(Value = "REBATE")]
            REBATE = 13,

            /// <summary>
            /// Enum GOVGRANT for value: GOV_GRANT
            /// </summary>
            [EnumMember(Value = "GOV_GRANT")]
            GOVGRANT = 14,

            /// <summary>
            /// Enum TAX for value: TAX
            /// </summary>
            [EnumMember(Value = "TAX")]
            TAX = 15,

            /// <summary>
            /// Enum FEE for value: FEE
            /// </summary>
            [EnumMember(Value = "FEE")]
            FEE = 16,

            /// <summary>
            /// Enum REI for value: REI
            /// </summary>
            [EnumMember(Value = "REI")]
            REI = 17,

            /// <summary>
            /// Enum FXT for value: FXT
            /// </summary>
            [EnumMember(Value = "FXT")]
            FXT = 18

        }


        /// <summary>
        /// Gets or Sets Type
        /// </summary>
        [DataMember(Name = "type", EmitDefaultValue = false)]
        public TypeEnum? Type { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="UniversalActivity" /> class.
        /// </summary>
        /// <param name="id">id.</param>
        /// <param name="account">account.</param>
        /// <param name="amount">amount.</param>
        /// <param name="currency">currency.</param>
        /// <param name="description">description.</param>
        /// <param name="fee">fee.</param>
        /// <param name="institution">institution.</param>
        /// <param name="optionType">If an option transaction, then it&#39;s type (BUY_TO_OPEN, SELL_TO_CLOSE, etc), otherwise empty string.</param>
        /// <param name="price">price.</param>
        /// <param name="settlementDate">settlementDate.</param>
        /// <param name="symbol">symbol.</param>
        /// <param name="optionSymbol">optionSymbol.</param>
        /// <param name="tradeDate">tradeDate.</param>
        /// <param name="type">type.</param>
        /// <param name="units">Usually but not necessarily an integer.</param>
        public UniversalActivity(string id = default(string), AccountSimple account = default(AccountSimple), decimal? amount = default(decimal?), Currency currency = default(Currency), string description = default(string), decimal fee = default(decimal), string institution = default(string), string optionType = default(string), decimal price = default(decimal), string settlementDate = default(string), Symbol symbol = default(Symbol), OptionsSymbol optionSymbol = default(OptionsSymbol), string tradeDate = default(string), TypeEnum? type = default(TypeEnum?), decimal units = default(decimal)) : base()
        {
            this.Id = id;
            this.Account = account;
            this.Amount = amount;
            this.Currency = currency;
            this.Description = description;
            this.Fee = fee;
            this.Institution = institution;
            this.OptionType = optionType;
            this.Price = price;
            this.SettlementDate = settlementDate;
            this.Symbol = symbol;
            this.OptionSymbol = optionSymbol;
            this.TradeDate = tradeDate;
            this.Type = type;
            this.Units = units;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// Gets or Sets Account
        /// </summary>
        [DataMember(Name = "account", EmitDefaultValue = false)]
        public AccountSimple Account { get; set; }

        /// <summary>
        /// Gets or Sets Amount
        /// </summary>
        [DataMember(Name = "amount", EmitDefaultValue = true)]
        public decimal? Amount { get; set; }

        /// <summary>
        /// Gets or Sets Currency
        /// </summary>
        [DataMember(Name = "currency", EmitDefaultValue = false)]
        public Currency Currency { get; set; }

        /// <summary>
        /// Gets or Sets Description
        /// </summary>
        [DataMember(Name = "description", EmitDefaultValue = false)]
        public string Description { get; set; }

        /// <summary>
        /// Gets or Sets Fee
        /// </summary>
        [DataMember(Name = "fee", EmitDefaultValue = false)]
        public decimal Fee { get; set; }

        /// <summary>
        /// Gets or Sets Institution
        /// </summary>
        [DataMember(Name = "institution", EmitDefaultValue = false)]
        public string Institution { get; set; }

        /// <summary>
        /// If an option transaction, then it&#39;s type (BUY_TO_OPEN, SELL_TO_CLOSE, etc), otherwise empty string
        /// </summary>
        /// <value>If an option transaction, then it&#39;s type (BUY_TO_OPEN, SELL_TO_CLOSE, etc), otherwise empty string</value>
        [DataMember(Name = "option_type", EmitDefaultValue = false)]
        public string OptionType { get; set; }

        /// <summary>
        /// Gets or Sets Price
        /// </summary>
        [DataMember(Name = "price", EmitDefaultValue = false)]
        public decimal Price { get; set; }

        /// <summary>
        /// Gets or Sets SettlementDate
        /// </summary>
        [DataMember(Name = "settlement_date", EmitDefaultValue = false)]
        public string SettlementDate { get; set; }

        /// <summary>
        /// Gets or Sets Symbol
        /// </summary>
        [DataMember(Name = "symbol", EmitDefaultValue = false)]
        public Symbol Symbol { get; set; }

        /// <summary>
        /// Gets or Sets OptionSymbol
        /// </summary>
        [DataMember(Name = "option_symbol", EmitDefaultValue = false)]
        public OptionsSymbol OptionSymbol { get; set; }

        /// <summary>
        /// Gets or Sets TradeDate
        /// </summary>
        [DataMember(Name = "trade_date", EmitDefaultValue = true)]
        public string TradeDate { get; set; }

        /// <summary>
        /// Usually but not necessarily an integer
        /// </summary>
        /// <value>Usually but not necessarily an integer</value>
        [DataMember(Name = "units", EmitDefaultValue = false)]
        public decimal Units { get; set; }

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
            sb.Append("class UniversalActivity {\n");
            sb.Append("  ").Append(base.ToString().Replace("\n", "\n  ")).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Account: ").Append(Account).Append("\n");
            sb.Append("  Amount: ").Append(Amount).Append("\n");
            sb.Append("  Currency: ").Append(Currency).Append("\n");
            sb.Append("  Description: ").Append(Description).Append("\n");
            sb.Append("  Fee: ").Append(Fee).Append("\n");
            sb.Append("  Institution: ").Append(Institution).Append("\n");
            sb.Append("  OptionType: ").Append(OptionType).Append("\n");
            sb.Append("  Price: ").Append(Price).Append("\n");
            sb.Append("  SettlementDate: ").Append(SettlementDate).Append("\n");
            sb.Append("  Symbol: ").Append(Symbol).Append("\n");
            sb.Append("  OptionSymbol: ").Append(OptionSymbol).Append("\n");
            sb.Append("  TradeDate: ").Append(TradeDate).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
            sb.Append("  Units: ").Append(Units).Append("\n");
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
            return this.Equals(input as UniversalActivity);
        }

        /// <summary>
        /// Returns true if UniversalActivity instances are equal
        /// </summary>
        /// <param name="input">Instance of UniversalActivity to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(UniversalActivity input)
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
                    this.Account == input.Account ||
                    (this.Account != null &&
                    this.Account.Equals(input.Account))
                ) && base.Equals(input) && 
                (
                    this.Amount == input.Amount ||
                    (this.Amount != null &&
                    this.Amount.Equals(input.Amount))
                ) && base.Equals(input) && 
                (
                    this.Currency == input.Currency ||
                    (this.Currency != null &&
                    this.Currency.Equals(input.Currency))
                ) && base.Equals(input) && 
                (
                    this.Description == input.Description ||
                    (this.Description != null &&
                    this.Description.Equals(input.Description))
                ) && base.Equals(input) && 
                (
                    this.Fee == input.Fee ||
                    this.Fee.Equals(input.Fee)
                ) && base.Equals(input) && 
                (
                    this.Institution == input.Institution ||
                    (this.Institution != null &&
                    this.Institution.Equals(input.Institution))
                ) && base.Equals(input) && 
                (
                    this.OptionType == input.OptionType ||
                    (this.OptionType != null &&
                    this.OptionType.Equals(input.OptionType))
                ) && base.Equals(input) && 
                (
                    this.Price == input.Price ||
                    this.Price.Equals(input.Price)
                ) && base.Equals(input) && 
                (
                    this.SettlementDate == input.SettlementDate ||
                    (this.SettlementDate != null &&
                    this.SettlementDate.Equals(input.SettlementDate))
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
                    this.TradeDate == input.TradeDate ||
                    (this.TradeDate != null &&
                    this.TradeDate.Equals(input.TradeDate))
                ) && base.Equals(input) && 
                (
                    this.Type == input.Type ||
                    this.Type.Equals(input.Type)
                ) && base.Equals(input) && 
                (
                    this.Units == input.Units ||
                    this.Units.Equals(input.Units)
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
                if (this.Account != null)
                {
                    hashCode = (hashCode * 59) + this.Account.GetHashCode();
                }
                if (this.Amount != null)
                {
                    hashCode = (hashCode * 59) + this.Amount.GetHashCode();
                }
                if (this.Currency != null)
                {
                    hashCode = (hashCode * 59) + this.Currency.GetHashCode();
                }
                if (this.Description != null)
                {
                    hashCode = (hashCode * 59) + this.Description.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Fee.GetHashCode();
                if (this.Institution != null)
                {
                    hashCode = (hashCode * 59) + this.Institution.GetHashCode();
                }
                if (this.OptionType != null)
                {
                    hashCode = (hashCode * 59) + this.OptionType.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Price.GetHashCode();
                if (this.SettlementDate != null)
                {
                    hashCode = (hashCode * 59) + this.SettlementDate.GetHashCode();
                }
                if (this.Symbol != null)
                {
                    hashCode = (hashCode * 59) + this.Symbol.GetHashCode();
                }
                if (this.OptionSymbol != null)
                {
                    hashCode = (hashCode * 59) + this.OptionSymbol.GetHashCode();
                }
                if (this.TradeDate != null)
                {
                    hashCode = (hashCode * 59) + this.TradeDate.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Type.GetHashCode();
                hashCode = (hashCode * 59) + this.Units.GetHashCode();
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
