<template>
  <v-container class="invoice-container">
    <div class="invoice">
      <!-- HEADER -->
      <div class="header">
        <div class="logo">
          <!-- Placeholder for Logo -->
          <div
            class="logo-placeholder"
            style="
              transform: scale(0.6);
              transform-origin: left center;
              width: 100px;
            "
          >
            <TestLogo />
          </div>
          <div>
            <strong>BAPUJI SURGICALS</strong><br />
            <small>Since 1988 in the service of humanity</small>
          </div>
        </div>

        <div class="company-details right">
          Ground Floor, 54/342, Chirakkaparambil<br />
          Kadavanthra, Kochi – 682020<br />
          Kerala, India<br />
          GSTIN: 32ANVPA2986J1ZW
        </div>
      </div>

      <div class="title">TAX INVOICE</div>

      <!-- RECIPIENT & INVOICE DETAILS -->
      <table>
        <tr>
          <th width="50%">Details of Recipient (Bill To Party)</th>
          <th width="50%">Invoice Details</th>
        </tr>
        <tr>
          <td class="input-cell">
            <v-combobox
              v-model="billTo.name"
              :items="hospitals"
              label="Recipient Name"
              density="compact"
              variant="underlined"
              hide-details
              @update:model-value="onHospitalChange"
            ></v-combobox>
            <v-textarea
              v-model="billTo.address"
              label="Address"
              rows="3"
              density="compact"
              variant="underlined"
              hide-details
              class="mt-1"
            ></v-textarea>
            <v-text-field
              v-model="billTo.gstin"
              label="GSTIN"
              density="compact"
              variant="underlined"
              hide-details
            ></v-text-field>
          </td>
          <td class="input-cell">
            <div class="d-flex align-center gap-2">
              <span class="label">Invoice No:</span>
              <v-text-field
                v-model="invoice.no"
                density="compact"
                variant="underlined"
                hide-details
              ></v-text-field>
            </div>
            <div class="d-flex align-center gap-2">
              <span class="label">Date:</span>
              <input type="date" v-model="invoice.date" class="simple-input" />
            </div>
            <div class="d-flex align-center gap-2">
              <span class="label">Rev. Charge:</span>
              <select v-model="invoice.reverseCharge" class="simple-input">
                <option :value="true">Yes</option>
                <option :value="false">No</option>
              </select>
            </div>
            <div class="d-flex align-center gap-2">
              <span class="label">Terms:</span>
              <v-text-field
                v-model="invoice.terms"
                density="compact"
                variant="underlined"
                hide-details
              ></v-text-field>
            </div>
          </td>
        </tr>
      </table>

      <!-- CONSIGNEE -->
      <table>
        <tr>
          <th class="d-flex justify-space-between align-center">
            Details of Consignee (Ship To Party)
            <v-btn size="x-small" variant="text" @click="copyBillToShip">
              Same as Bill To
            </v-btn>
          </th>
        </tr>
        <tr>
          <td class="input-cell">
            <v-text-field
              v-model="shipTo.name"
              label="Consignee Name"
              density="compact"
              variant="underlined"
              hide-details
            ></v-text-field>
            <v-textarea
              v-model="shipTo.address"
              label="Address"
              rows="3"
              density="compact"
              variant="underlined"
              hide-details
              class="mt-1"
            ></v-textarea>
            <v-text-field
              v-model="shipTo.gstin"
              label="GSTIN"
              density="compact"
              variant="underlined"
              hide-details
            ></v-text-field>
          </td>
        </tr>
      </table>

      <!-- ITEMS TABLE -->
      <table class="items">
        <thead>
          <tr>
            <th width="5%">Sl No</th>
            <th width="40%">Description</th>
            <th width="10%">HSN</th>
            <th width="10%">Qty</th>
            <th width="15%">Rate</th>
            <th width="15%">Amount</th>
            <th width="5%" class="no-print"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in items" :key="index">
            <td align="center">{{ index + 1 }}</td>
            <td>
              <v-text-field
                v-model="item.description"
                density="compact"
                variant="plain"
                hide-details
                placeholder="Item Name"
              ></v-text-field>
            </td>
            <td align="center">
              <v-text-field
                v-model="item.hsn"
                density="compact"
                variant="plain"
                hide-details
                placeholder="HSN"
              ></v-text-field>
            </td>
            <td align="center">
              <v-text-field
                v-model.number="item.qty"
                type="number"
                density="compact"
                variant="plain"
                hide-details
                min="1"
              ></v-text-field>
            </td>
            <td class="right">
              <v-text-field
                v-model.number="item.rate"
                type="number"
                density="compact"
                variant="plain"
                hide-details
                class="right-input"
                min="0"
                step="0.01"
              ></v-text-field>
            </td>
            <td class="right">{{ formatCurrency(item.qty * item.rate) }}</td>
            <td class="no-print">
              <v-icon size="small" color="error" @click="removeItem(index)"
                >mdi-delete</v-icon
              >
            </td>
          </tr>
        </tbody>
        <tfoot>
          <tr class="no-print">
            <td colspan="7" class="text-center">
              <v-btn block variant="tonal" size="small" @click="addItem"
                >+ Add Item</v-btn
              >
            </td>
          </tr>
          <tr>
            <td colspan="5" class="right bold">Total</td>
            <td class="right bold">{{ formatCurrency(subTotal) }}</td>
            <td></td>
          </tr>
        </tfoot>
      </table>

      <!-- BANK & TAX SUMMARY -->
      <table>
        <tr>
          <td width="50%">
            <strong>A/c No</strong> : 424600860000011<br />
            <strong>Bank Name</strong> : Punjab National Bank<br />
            <strong>Branch</strong> : Lal Bagh, Fort Road, Bangalore<br />
            <strong>IFSC Code</strong> : PUNB0424600
          </td>
          <td width="50%">
            <table class="no-border">
              <tr>
                <td>Add Output CGST ({{ cgstRate }}%)</td>
                <td class="right">{{ formatCurrency(taxAmount / 2) }}</td>
              </tr>
              <tr>
                <td>Add Output SGST ({{ sgstRate }}%)</td>
                <td class="right">{{ formatCurrency(taxAmount / 2) }}</td>
              </tr>
              <tr>
                <td>Add Output IGST</td>
                <td class="right">0.00</td>
              </tr>
              <tr class="bold">
                <td>Tax Amount GST</td>
                <td class="right">{{ formatCurrency(taxAmount) }}</td>
              </tr>
              <tr>
                <td>Round Off</td>
                <td class="right">{{ roundOff }}</td>
              </tr>
              <tr class="bold">
                <td>Total Amount After Tax</td>
                <td class="right">{{ formatCurrency(grandTotal) }}</td>
              </tr>
            </table>
          </td>
        </tr>
      </table>

      <!-- AMOUNT IN WORDS -->
      <table>
        <tr>
          <td class="bold">
            Total Invoice Amount (Rupees in Words) :
            <span style="font-weight: normal"> {{ amountInWords }} only </span>
          </td>
        </tr>
      </table>

      <!-- DECLARATION -->
      <table>
        <tr>
          <td width="65%">
            <strong>Declaration</strong><br /><br />
            1. Goods once sold will not be taken back.<br />
            2. Interest 24% P.A. on overdue accounts.<br />
            3. Errors & omissions excepted.<br />
            4. Subject to Bangalore Jurisdiction.
          </td>
          <td width="35%" class="right">
            <strong>For BAPUJI SURGICALS</strong><br /><br />
            Joseph
            <div class="sign-box"></div>
            <strong>AUTHORISED SIGNATORY</strong>
          </td>
        </tr>
      </table>

      <!-- FOOTER -->
      <div class="footer-line">
        <div>Date & Time : {{ currentDateTime }}</div>
        <div>PH: 8870472323 | info@bapujisurgicals.com</div>
        <div>Page 1 of 1</div>
      </div>
    </div>
  </v-container>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import TestLogo from "@/components/TestLogo.vue";

export default {
  name: "BillingPage",
  components: {
    TestLogo,
  },
  data() {
    return {
      billTo: { name: "", address: "", gstin: "" },
      shipTo: { name: "", address: "", gstin: "" },
      invoice: {
        no: "25321438",
        date: new Date().toISOString().substr(0, 10),
        reverseCharge: false,
        terms: "60 Days",
      },
      items: [
        {
          description: "Surgical Item Name",
          hsn: "3004",
          qty: 10,
          rate: 500.0,
        },
      ],
      cgstRate: 9,
      sgstRate: 9,
    };
  },
  computed: {
    ...mapGetters({
      hospitals: "getHospitals",
    }),
    subTotal() {
      return this.items.reduce((sum, item) => sum + item.qty * item.rate, 0);
    },
    taxAmount() {
      // Assuming 18% total tax for now based on template logic splitting into CGST/SGST
      return this.subTotal * ((this.cgstRate + this.sgstRate) / 100);
    },
    totalWithTax() {
      return this.subTotal + this.taxAmount;
    },
    grandTotal() {
      return Math.round(this.totalWithTax);
    },
    roundOff() {
      return (this.grandTotal - this.totalWithTax).toFixed(2);
    },
    currentDateTime() {
      return new Date().toLocaleString();
    },
    amountInWords() {
      return this.numberToWords(this.grandTotal);
    },
  },
  methods: {
    ...mapActions(["addHospital"]),
    onHospitalChange(val) {
      if (val && !this.hospitals.includes(val)) {
        this.addHospital(val);
      }
      this.billTo.name = val;
    },
    addItem() {
      this.items.push({ description: "", hsn: "", qty: 1, rate: 0 });
    },
    removeItem(index) {
      this.items.splice(index, 1);
    },
    copyBillToShip() {
      this.shipTo = { ...this.billTo };
    },
    formatCurrency(value) {
      if (!value) return "0.00";
      return value.toLocaleString("en-IN", {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2,
      });
    },
    numberToWords(number) {
      // Simple basic implementation for Indian Numbering System
      const a = [
        "",
        "One ",
        "Two ",
        "Three ",
        "Four ",
        "Five ",
        "Six ",
        "Seven ",
        "Eight ",
        "Nine ",
        "Ten ",
        "Eleven ",
        "Twelve ",
        "Thirteen ",
        "Fourteen ",
        "Fifteen ",
        "Sixteen ",
        "Seventeen ",
        "Eighteen ",
        "Nineteen ",
      ];
      const b = [
        "",
        "",
        "Twenty",
        "Thirty",
        "Forty",
        "Fifty",
        "Sixty",
        "Seventy",
        "Eighty",
        "Ninety",
      ];

      let num = parseFloat(number).toFixed(0);
      if ((num = num.toString()).length > 9) return "overflow";
      const n = ("000000000" + num)
        .substr(-9)
        .match(/^(\d{2})(\d{2})(\d{2})(\d{1})(\d{2})$/);
      if (!n) return;

      let str = "";
      str +=
        n[1] != 0
          ? (a[Number(n[1])] || b[n[1][0]] + " " + a[n[1][1]]) + "Crore "
          : "";
      str +=
        n[2] != 0
          ? (a[Number(n[2])] || b[n[2][0]] + " " + a[n[2][1]]) + "Lakh "
          : "";
      str +=
        n[3] != 0
          ? (a[Number(n[3])] || b[n[3][0]] + " " + a[n[3][1]]) + "Thousand "
          : "";
      str +=
        n[4] != 0
          ? (a[Number(n[4])] || b[n[4][0]] + " " + a[n[4][1]]) + "Hundred "
          : "";
      str +=
        n[5] != 0
          ? (str != "" ? "and " : "") +
            (a[Number(n[5])] || b[n[5][0]] + " " + a[n[5][1]])
          : "";

      return str || "Zero";
    },
  },
};
</script>

<style scoped>
/* Inheriting fonts/styles from the template provided */
.invoice-container {
  background: #f5f5f5;
  padding: 20px;
  font-family: Arial, Helvetica, sans-serif;
}

.invoice {
  width: 100%;
  max-width: 900px;
  /* Slight increase for editing comfort */
  margin: 0 auto;
  background: #fff;
  padding: 20px;
  border: 1px solid #000;
  font-size: 12px;
}

/* Header */
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 2px solid #000;
  padding-bottom: 10px;
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
}

.company-details {
  font-size: 12px;
  line-height: 1.4;
  text-align: right;
}

.title {
  text-align: center;
  font-weight: bold;
  font-size: 18px;
  margin: 10px 0;
}

/* Tables */
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}

th,
td {
  border: 1px solid #000;
  padding: 6px;
  vertical-align: top;
}

th {
  background: #eee;
  text-align: left;
  font-weight: bold;
}

.items th {
  text-align: center;
}

.right {
  text-align: right;
}

.bold {
  font-weight: bold;
}

.no-border,
.no-border tr,
.no-border td {
  border: none !important;
}

/* Form Inputs within Table */
.input-cell {
  padding: 4px 8px;
}

.gap-2 {
  gap: 8px;
}

.label {
  white-space: nowrap;
  font-weight: bold;
}

.simple-input {
  border-bottom: 1px solid #999;
  padding: 2px 5px;
  outline: none;
  width: 100%;
}

.right-input input {
  text-align: right;
}

/* Signature & footer */
.sign-box {
  height: 80px;
}

.footer-line {
  font-size: 11px;
  margin-top: 10px;
  display: flex;
  justify-content: space-between;
  border-top: 1px solid #000;
  padding-top: 5px;
}

/* Utility to hide elements during print if needed, though scoped css might be tricky with print */
@media print {
  .no-print {
    display: none !important;
  }

  .invoice-container {
    padding: 0;
    background: white;
  }

  .invoice {
    border: none;
    margin: 0;
    width: 100%;
    max-width: none;
  }
}
</style>
