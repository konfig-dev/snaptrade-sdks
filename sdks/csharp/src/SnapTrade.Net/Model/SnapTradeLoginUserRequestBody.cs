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
    /// Data to login a user via SnapTrade Partner
    /// </summary>
    [DataContract(Name = "SnapTradeLoginUserRequestBody")]
    public partial class SnapTradeLoginUserRequestBody : IEquatable<SnapTradeLoginUserRequestBody>, IValidatableObject
    {
        /// <summary>
        /// Sets whether the connection should be read-only or trade-enabled. Defaults to read-only if not specified.
        /// </summary>
        /// <value>Sets whether the connection should be read-only or trade-enabled. Defaults to read-only if not specified.</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum ConnectionTypeEnum
        {
            /// <summary>
            /// Enum Read for value: read
            /// </summary>
            [EnumMember(Value = "read")]
            Read = 1,

            /// <summary>
            /// Enum Trade for value: trade
            /// </summary>
            [EnumMember(Value = "trade")]
            Trade = 2

        }


        /// <summary>
        /// Sets whether the connection should be read-only or trade-enabled. Defaults to read-only if not specified.
        /// </summary>
        /// <value>Sets whether the connection should be read-only or trade-enabled. Defaults to read-only if not specified.</value>
        [DataMember(Name = "connectionType", EmitDefaultValue = false)]
        public ConnectionTypeEnum? ConnectionType { get; set; }
        /// <summary>
        /// Sets the connection portal version to render. Currently only v4 is supported and is the default. All other versions are deprecated and will automatically be set to v4.
        /// </summary>
        /// <value>Sets the connection portal version to render. Currently only v4 is supported and is the default. All other versions are deprecated and will automatically be set to v4.</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum ConnectionPortalVersionEnum
        {
            /// <summary>
            /// Enum V4 for value: v4
            /// </summary>
            [EnumMember(Value = "v4")]
            V4 = 1,

            /// <summary>
            /// Enum V3 for value: v3
            /// </summary>
            [EnumMember(Value = "v3")]
            V3 = 2,

            /// <summary>
            /// Enum V2 for value: v2
            /// </summary>
            [EnumMember(Value = "v2")]
            V2 = 3

        }


        /// <summary>
        /// Sets the connection portal version to render. Currently only v4 is supported and is the default. All other versions are deprecated and will automatically be set to v4.
        /// </summary>
        /// <value>Sets the connection portal version to render. Currently only v4 is supported and is the default. All other versions are deprecated and will automatically be set to v4.</value>
        [DataMember(Name = "connectionPortalVersion", EmitDefaultValue = false)]
        public ConnectionPortalVersionEnum? ConnectionPortalVersion { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="SnapTradeLoginUserRequestBody" /> class.
        /// </summary>
        /// <param name="broker">Slug of the brokerage to connect the user to. See [the integrations page](https://snaptrade.notion.site/66793431ad0b416489eaabaf248d0afb?v&#x3D;3cfea70ef4254afc89704e47275a7a9a&amp;pvs&#x3D;4) for a list of supported brokerages and their slugs..</param>
        /// <param name="immediateRedirect">When set to &#x60;true&#x60;, user will be redirected back to the partner&#39;s site instead of the connection portal. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information..</param>
        /// <param name="customRedirect">URL to redirect the user to after the user connects their brokerage account. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information..</param>
        /// <param name="reconnect">The UUID of the brokerage connection to be reconnected. This parameter should be left empty unless you are reconnecting a disabled connection. See the [guide on fixing broken connections](/docs/fix-broken-connections) for more information..</param>
        /// <param name="connectionType">Sets whether the connection should be read-only or trade-enabled. Defaults to read-only if not specified. (default to ConnectionTypeEnum.Read).</param>
        /// <param name="connectionPortalVersion">Sets the connection portal version to render. Currently only v4 is supported and is the default. All other versions are deprecated and will automatically be set to v4. (default to ConnectionPortalVersionEnum.V4).</param>
        public SnapTradeLoginUserRequestBody(string broker = default(string), bool immediateRedirect = default(bool), string customRedirect = default(string), string reconnect = default(string), ConnectionTypeEnum? connectionType = ConnectionTypeEnum.Read, ConnectionPortalVersionEnum? connectionPortalVersion = ConnectionPortalVersionEnum.V4)
        {
            this.Broker = broker;
            this.ImmediateRedirect = immediateRedirect;
            this.CustomRedirect = customRedirect;
            this.Reconnect = reconnect;
            this.ConnectionType = connectionType;
            this.ConnectionPortalVersion = connectionPortalVersion;
        }

        /// <summary>
        /// Slug of the brokerage to connect the user to. See [the integrations page](https://snaptrade.notion.site/66793431ad0b416489eaabaf248d0afb?v&#x3D;3cfea70ef4254afc89704e47275a7a9a&amp;pvs&#x3D;4) for a list of supported brokerages and their slugs.
        /// </summary>
        /// <value>Slug of the brokerage to connect the user to. See [the integrations page](https://snaptrade.notion.site/66793431ad0b416489eaabaf248d0afb?v&#x3D;3cfea70ef4254afc89704e47275a7a9a&amp;pvs&#x3D;4) for a list of supported brokerages and their slugs.</value>
        [DataMember(Name = "broker", EmitDefaultValue = false)]
        public string Broker { get; set; }

        /// <summary>
        /// When set to &#x60;true&#x60;, user will be redirected back to the partner&#39;s site instead of the connection portal. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information.
        /// </summary>
        /// <value>When set to &#x60;true&#x60;, user will be redirected back to the partner&#39;s site instead of the connection portal. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information.</value>
        [DataMember(Name = "immediateRedirect", EmitDefaultValue = true)]
        public bool ImmediateRedirect { get; set; }

        /// <summary>
        /// URL to redirect the user to after the user connects their brokerage account. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information.
        /// </summary>
        /// <value>URL to redirect the user to after the user connects their brokerage account. This parameter is ignored if the connection portal is loaded inside an iframe. See the [guide on ways to integrate the connection portal](/docs/implement-connection-portal) for more information.</value>
        [DataMember(Name = "customRedirect", EmitDefaultValue = false)]
        public string CustomRedirect { get; set; }

        /// <summary>
        /// The UUID of the brokerage connection to be reconnected. This parameter should be left empty unless you are reconnecting a disabled connection. See the [guide on fixing broken connections](/docs/fix-broken-connections) for more information.
        /// </summary>
        /// <value>The UUID of the brokerage connection to be reconnected. This parameter should be left empty unless you are reconnecting a disabled connection. See the [guide on fixing broken connections](/docs/fix-broken-connections) for more information.</value>
        [DataMember(Name = "reconnect", EmitDefaultValue = false)]
        public string Reconnect { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class SnapTradeLoginUserRequestBody {\n");
            sb.Append("  Broker: ").Append(Broker).Append("\n");
            sb.Append("  ImmediateRedirect: ").Append(ImmediateRedirect).Append("\n");
            sb.Append("  CustomRedirect: ").Append(CustomRedirect).Append("\n");
            sb.Append("  Reconnect: ").Append(Reconnect).Append("\n");
            sb.Append("  ConnectionType: ").Append(ConnectionType).Append("\n");
            sb.Append("  ConnectionPortalVersion: ").Append(ConnectionPortalVersion).Append("\n");
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
            return this.Equals(input as SnapTradeLoginUserRequestBody);
        }

        /// <summary>
        /// Returns true if SnapTradeLoginUserRequestBody instances are equal
        /// </summary>
        /// <param name="input">Instance of SnapTradeLoginUserRequestBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(SnapTradeLoginUserRequestBody input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Broker == input.Broker ||
                    (this.Broker != null &&
                    this.Broker.Equals(input.Broker))
                ) && 
                (
                    this.ImmediateRedirect == input.ImmediateRedirect ||
                    this.ImmediateRedirect.Equals(input.ImmediateRedirect)
                ) && 
                (
                    this.CustomRedirect == input.CustomRedirect ||
                    (this.CustomRedirect != null &&
                    this.CustomRedirect.Equals(input.CustomRedirect))
                ) && 
                (
                    this.Reconnect == input.Reconnect ||
                    (this.Reconnect != null &&
                    this.Reconnect.Equals(input.Reconnect))
                ) && 
                (
                    this.ConnectionType == input.ConnectionType ||
                    this.ConnectionType.Equals(input.ConnectionType)
                ) && 
                (
                    this.ConnectionPortalVersion == input.ConnectionPortalVersion ||
                    this.ConnectionPortalVersion.Equals(input.ConnectionPortalVersion)
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
                if (this.Broker != null)
                {
                    hashCode = (hashCode * 59) + this.Broker.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.ImmediateRedirect.GetHashCode();
                if (this.CustomRedirect != null)
                {
                    hashCode = (hashCode * 59) + this.CustomRedirect.GetHashCode();
                }
                if (this.Reconnect != null)
                {
                    hashCode = (hashCode * 59) + this.Reconnect.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.ConnectionType.GetHashCode();
                hashCode = (hashCode * 59) + this.ConnectionPortalVersion.GetHashCode();
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
