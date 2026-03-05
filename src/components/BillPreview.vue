<template>
  <div class="invoice-container">
    <div class="main-header-title">TAX INVOICE</div>
    <div class="invoice-box">
      <!-- HEADER -->
      <div class="header-section border-bottom">
        <div class="company-name text-center">GOODWILL ENTERPRISES</div>
        <div class="company-address text-center">
          D.N.PP/2/175 Kottara, Meeyannoor P.O. Kollam - 691537<br />
          Ph. 9946719904, 8281664274
        </div>

        <div class="header-details d-flex justify-space-between mt-2">
          <div class="left-details">
            <div>
              <strong>GSTIN:</strong>
              {{ companyProfile.gstin || "32BJMPM1863M1ZU" }}
            </div>
            <div>
              <strong>Reg. No.</strong>
              {{ companyProfile.regNo || "MDKER02220001" }}
            </div>
          </div>
          <div class="right-details text-right">
            <div>
              <strong>DLNO.</strong>
              {{ companyProfile.dlNo1 || "20 B WLF20B2025KL001784" }}
            </div>
            <div>
              <strong>21 B</strong>
              {{ companyProfile.dlNo2 || "WLF21B2025KL001750" }}
            </div>
            <div>
              <strong>Email.</strong>
              {{ companyProfile.email || "gektra19@gmail.com" }}
            </div>
          </div>
        </div>
      </div>

      <!-- BUYER TO / INVOICE DETAILS -->
      <div class="buyer-section border-bottom d-flex">
        <div class="buyer-details border-right" style="width: 60%">
          <div class="section-label">Buyer(Bill to)</div>
          <div class="buyer-name">
            <strong>{{ recipient.name || "N/A" }}</strong>
            <div v-if="partyName" class="text-caption">
              Party: <b>{{ partyName }}</b>
            </div>
          </div>
          <div class="buyer-address">
            {{ recipient.address || "" }}<br />
            {{ recipient.city
            }}{{ recipient.state ? ", " + recipient.state : "" }}
          </div>
          <div class="buyer-meta mt-1">
            <div><strong>DLNO.</strong> {{ recipient.dlNo || "N/A" }}</div>
            <div><strong>21 B</strong> {{ recipient.dlNo2 || "N/A" }}</div>
            <div><strong>GSTIN :</strong> {{ recipient.gstin || "N/A" }}</div>
          </div>
        </div>
        <div class="invoice-meta" style="width: 40%; padding-left: 10px">
          <div class="d-flex justify-space-between">
            <span>Invoice No . {{ invoice.gstInvoiceNo || "H.40" }}</span>
          </div>
          <div class="d-flex justify-space-between">
            <span
              >Date .
              {{
                invoice.gstInvoiceDate || new Date().toLocaleDateString()
              }}</span
            >
          </div>
        </div>
      </div>

      <!-- ITEMS TABLE -->
      <div class="items-section border-bottom">
        <table class="custom-table">
          <thead>
            <tr>
              <th style="width: 5%">Sl.No</th>
              <th style="width: 35%">Description of Goods</th>
              <th style="width: 10%">HSN/SAC</th>
              <th style="width: 8%">Quantity</th>
              <th style="width: 10%">Rate</th>
              <th style="width: 10%">MRP</th>
              <th style="width: 8%">GST %</th>
              <th style="width: 14%">Amount</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in items" :key="index">
              <td class="text-center" style="vertical-align: top">
                {{ index + 1 }}
              </td>
              <td style="vertical-align: top">
                <strong>{{ item.medicineName }}</strong
                ><br />
                <div class="item-meta" style="font-size: 0.9em">
                  Batch . &nbsp; {{ item.batchId }}<br />
                  Mfg Dt. &nbsp; {{ item.mfgDate || "Jun-2025" }}<br />
                  Expiry. &nbsp; {{ item.expiryDate }}
                </div>
              </td>
              <td class="text-center" style="vertical-align: top">
                {{ item.hsnCode || "90183290" }}
              </td>
              <td class="text-center" style="vertical-align: top">
                {{ item.quantity }}
              </td>
              <td class="text-right" style="vertical-align: top">
                {{ parseFloat(item.rate || item.mrp).toFixed(2) }}
              </td>
              <td class="text-right" style="vertical-align: top">
                {{ parseFloat(item.mrp).toFixed(2) }}
              </td>
              <td class="text-center" style="vertical-align: top">
                {{ item.gst }}%
              </td>
              <td class="text-right" style="vertical-align: top">
                <strong>{{ calculateAmount(item).toFixed(2) }}</strong>
              </td>
            </tr>
            <!-- Spacer rows to fill height if needed -->
            <tr style="height: 100px; border: none">
              <td colspan="8" style="vertical-align: bottom; border: none">
                <div class="d-flex justify-end p-2" v-if="totalTax > 0">
                  <div
                    class="tax-summary text-right"
                    style="margin-right: 20px"
                  >
                    <div>Output CGST @ {{ items[0]?.gst / 2 || 2.5 }}%</div>
                    <div>Output SGST @ {{ items[0]?.gst / 2 || 2.5 }}%</div>
                    <div><i>Round Off</i></div>
                  </div>
                  <div class="tax-values text-right">
                    <div>{{ (totalTax / 2).toFixed(2) }}</div>
                    <div>{{ (totalTax / 2).toFixed(2) }}</div>
                    <div>0.06</div>
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
          <tfoot>
            <tr style="border-top: 2px solid black">
              <td colspan="3" class="text-right"><strong>Total</strong></td>
              <td class="text-center">
                <strong>{{ totalQuantity }}</strong>
              </td>
              <td></td>
              <td></td>
              <td></td>
              <td class="text-right">
                <strong>{{ grandTotal.toFixed(2) }}</strong>
              </td>
            </tr>
          </tfoot>
        </table>
      </div>

      <!-- AMOUNT IN WORDS -->
      <div class="amount-words-section border-bottom p-2">
        <div>Amount Chargeable (in words )</div>
        <div style="font-weight: bold; margin-top: 5px">
          INR {{ amountInWords }} Only.
        </div>
      </div>

      <!-- TAX BREAKDOWN TABLE -->
      <div class="tax-breakdown-section border-bottom">
        <table class="custom-table">
          <thead>
            <tr>
              <th style="width: 25%">HSN/SAC</th>
              <th style="width: 20%">Taxable Value</th>
              <th colspan="2" style="width: 25%">CGST</th>
              <th colspan="2" style="width: 25%">SGST</th>
              <th style="width: 5%">Total</th>
            </tr>
            <tr>
              <th></th>
              <th></th>
              <th>Rate</th>
              <th>Amount</th>
              <th>Rate</th>
              <th>Amount</th>
              <th>Tax Amount</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(taxItem, idx) in taxSummary" :key="idx">
              <td class="text-center">{{ taxItem.hsn }}</td>
              <td class="text-right">{{ taxItem.taxable.toFixed(2) }}</td>
              <td class="text-center">{{ taxItem.gstRate / 2 }}%</td>
              <td class="text-right">{{ (taxItem.taxAmt / 2).toFixed(2) }}</td>
              <td class="text-center">{{ taxItem.gstRate / 2 }}%</td>
              <td class="text-right">{{ (taxItem.taxAmt / 2).toFixed(2) }}</td>
              <td class="text-right">{{ taxItem.taxAmt.toFixed(2) }}</td>
            </tr>
            <tr style="font-weight: bold">
              <td class="text-left">Total</td>
              <td class="text-right">{{ (subTotal - totalTax).toFixed(2) }}</td>
              <td></td>
              <td class="text-right">{{ (totalTax / 2).toFixed(2) }}</td>
              <td></td>
              <td class="text-right">{{ (totalTax / 2).toFixed(2) }}</td>
              <td class="text-right">{{ totalTax.toFixed(2) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- BANK & Footer -->
      <div class="footer-section d-flex">
        <div
          class="declaration-section border-right"
          style="width: 50%; padding: 5px"
        >
          <strong>DECLARATION</strong><br />
          <div style="font-size: 0.9em; margin-top: 5px">
            We declare that this invoice shows the actual price of the goods
            described and that all particulars are true and correct .
          </div>
          <div class="mt-4">Customer Signature</div>
        </div>
        <div class="bank-section" style="width: 50%; padding: 5px">
          <div class="mb-2"><strong>Bank Details</strong></div>
          <div>Bank Name : {{ bank.bankName || "FEDERAL BANK" }}</div>
          <div>A/c No. : {{ bank.accountNumber || "17280200001260" }}</div>
          <div>Branch & IFSCode : {{ bank.branch }} & {{ bank.ifscCode }}</div>

          <div class="text-right mt-4">
            For &nbsp; <strong>GOODWILL ENTERPRISES</strong> <br /><br /><br />
            Authorised Signatory
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "BillPreview",
  props: {
    recipient: { type: Object, default: () => ({}) },
    consignee: { type: Object, default: () => ({}) },
    invoice: { type: Object, default: () => ({}) },
    bank: { type: Object, default: () => ({}) },
    items: { type: Array, default: () => [] },
    subTotal: { type: Number, default: 0 },
    discount: { type: Object, default: () => ({ value: 0, amount: 0 }) },
    grandTotal: { type: Number, default: 0 },
    partyName: { type: String, default: "" },
  },
  computed: {
    companyProfile() {
      // Mock company details - in a real app, pass this as a prop
      return {
        name: "GOODWILL ENTERPRISES",
        address1: "D.N.PP/2/175 Kottara, Meeyannoor P.O. Kollam - 691537",
        phone: "9946719904 ,8281664274",
        gstin: "32BJMPM1863M1ZU",
        regNo: "MDKER02220001",
        dlNo1: "20 B WLF20B2025KL001784",
        dlNo2: "21 B WLF21B2025KL001750",
        email: "gektra19@gmail.com",
      };
    },
    totalQuantity() {
      return this.items.reduce(
        (acc, item) => acc + (parseInt(item.quantity) || 0),
        0
      );
    },
    totalTax() {
      return this.items.reduce((acc, item) => {
        const qty = parseFloat(item.quantity) || 0;
        const mrp = parseFloat(item.mrp) || 0;
        const gst = parseFloat(item.gst) || 0;
        // Logic: Rate * Qty = Taxable? Or MRP?
        // Assuming Amount = Taxable + Tax
        // For simple display, let's assume MRP is inclusive or calc tax on base.
        // Let's use the calculateAmount logic
        const baseAmount = mrp * qty; // Using MRP as base for now, can be Rate
        const gstAmount = baseAmount * (gst / 100);
        return acc + gstAmount;
      }, 0);
    },
    taxSummary() {
      // Group items by HSN and GST Rate to show summary
      const summary = {};
      this.items.forEach((item) => {
        const key = `${item.hsnCode || "90183290"}-${item.gst}`;
        if (!summary[key]) {
          summary[key] = {
            hsn: item.hsnCode || "90183290",
            gstRate: parseFloat(item.gst),
            taxable: 0,
            taxAmt: 0,
          };
        }
        const qty = parseFloat(item.quantity) || 0;
        const rate = parseFloat(item.rate || item.mrp) || 0; // Taxable Rate

        const taxableVal = rate * qty;
        const taxVal = taxableVal * (parseFloat(item.gst) / 100);

        summary[key].taxable += taxableVal;
        summary[key].taxAmt += taxVal;
      });
      return Object.values(summary);
    },
    amountInWords() {
      const num = Math.round(this.grandTotal);
      return this.numberToEnglish(num);
    },
  },
  methods: {
    calculateAmount(item) {
      const qty = parseFloat(item.quantity) || 0;
      const rate = parseFloat(item.rate || item.mrp) || 0; // Use Rate if available, else MRP
      // If we assume MRP is inclusive, logic differs.
      // User image shows: Rate=1702.80, MRP=3960.00, GST=5%, Amount=1702.80.
      // Wait, 1702.80 * 1 = 1702.80.
      // Below: Taxable Value = 1702.80. CGST 2.5% = 42.57.
      // So Amount column in main table seems to be Taxable Value.
      // Grand Total includes Tax.
      return rate * qty; // Main table amount usually corresponds to taxable value in this format
    },
    numberToEnglish(n) {
      if (n < 0) return "Negative";
      if (n === 0) return "Zero";

      var string = n.toString(),
        units,
        tens,
        scales,
        start,
        end,
        chunks,
        chunksLen,
        chunk,
        ints,
        i,
        word,
        words;
      var and = chunksLen > 0 ? "and" : "";

      /* Is number zero? */
      if (parseInt(string) === 0) {
        return "Zero";
      }

      /* Array of units as words */
      units = [
        "",
        "One",
        "Two",
        "Three",
        "Four",
        "Five",
        "Six",
        "Seven",
        "Eight",
        "Nine",
        "Ten",
        "Eleven",
        "Twelve",
        "Thirteen",
        "Fourteen",
        "Fifteen",
        "Sixteen",
        "Seventeen",
        "Eighteen",
        "Nineteen",
      ];

      /* Array of tens as words */
      tens = [
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

      /* Array of scales as words */
      scales = [
        "",
        "Thousand",
        "Million",
        "Billion",
        "Trillion",
        "Quadrillion",
        "Quintillion",
        "Sextillion",
        "Septillion",
        "Octillion",
        "Nonillion",
        "Decillion",
        "Undecillion",
        "Duodecillion",
        "Tredecillion",
        "Quattuordecillion",
        "Quindecillion",
        "Sexdecillion",
        "Septendecillion",
        "Octodecillion",
        "Novemdecillion",
        "Vigintillion",
      ];

      /* Split user argument into 3 digit chunks from right to left */
      start = string.length;
      chunks = [];
      while (start > 0) {
        end = start;
        chunks.push(string.slice((start = Math.max(0, start - 3)), end));
      }

      /* Check if function has enough scale words to be able to stringify the user argument */
      chunksLen = chunks.length;
      if (chunksLen > scales.length) {
        return "";
      }

      /* Stringify each integer in each three digit chunk */
      words = [];
      for (i = 0; i < chunksLen; i++) {
        chunk = parseInt(chunks[i]);
        if (chunk) {
          /* Split chunk into array of individual integers */
          ints = chunks[i].split("").reverse().map(parseFloat);
          /* If tens integer is 1, i.e. 10, then add 10 to units integer */
          if (ints[1] === 1) {
            ints[0] += 10;
          }
          /* Add scale word if chunk is not zero and array item exists */
          if ((word = scales[i])) {
            words.push(word);
          }
          /* Add unit word if array item exists */
          if ((word = units[ints[0]])) {
            words.push(word);
          }
          /* Add tens word if array item exists */
          if ((word = tens[ints[1]])) {
            words.push(word);
          }
          /* Add 'and' string after units or tens integer if: */
          if (ints[0] || ints[1]) {
            /* Chunk has a hundreds integer or chunk is the first of multiple chunks */
            if (ints[2] || (!i && chunksLen)) {
              words.push(and);
            }
          }
          /* Add hundreds word if array item exists */
          if ((word = units[ints[2]])) {
            words.push(word + " Hundred");
          }
        }
      }
      return words.reverse().join(" ") + " Only";
    },
  },
};
</script>

<style scoped>
.invoice-container {
  font-family: Arial, Helvetica, sans-serif;
  background: #f5f5f5;
  padding: 20px;
  overflow: auto;
  color: #000;
}

.main-header-title {
  text-align: center;
  font-weight: bold;
  font-size: 14px;
  margin-bottom: 2px;
}

.invoice-box {
  width: 100%;
  max-width: 900px;
  margin: 0 auto;
  background: #fff;
  border: 1px solid #000;
  box-sizing: border-box;
  font-size: 12px;
}

/* Utilities */
.text-center {
  text-align: center;
}

.text-right {
  text-align: right;
}

.text-left {
  text-align: left;
}

.bold {
  font-weight: bold;
}

.d-flex {
  display: flex;
}

.justify-space-between {
  justify-content: space-between;
}

.justify-end {
  justify-content: flex-end;
}

.mt-2 {
  margin-top: 8px;
}

.mt-4 {
  margin-top: 16px;
}

.mb-2 {
  margin-bottom: 8px;
}

.p-2 {
  padding: 8px;
}

/* Borders */
.border-bottom {
  border-bottom: 1px solid #000;
}

.border-right {
  border-right: 1px solid #000;
}

/* Header */
.header-section {
  padding: 10px;
}

.company-name {
  font-size: 20px;
  font-weight: 900;
  text-transform: uppercase;
}

.company-address {
  font-size: 11px;
  font-style: italic;
  font-weight: bold;
  margin-top: 5px;
}

.header-details {
  font-size: 11px;
  line-height: 1.4;
}

/* Buyer */
.buyer-section {
  display: flex;
}

.buyer-details,
.invoice-meta {
  padding: 8px;
}

.section-label {
  font-size: 10px;
}

.custom-table {
  width: 100%;
  border-collapse: collapse;
}

.custom-table th {
  border: 1px solid #000;
  border-top: none;
  padding: 5px;
  background: #fff;
  font-weight: bold;
  text-align: center;
}

.custom-table td {
  border-left: 1px solid #000;
  border-right: 1px solid #000;
  padding: 5px;
}

/* .custom-table tbody tr:last-child td { } */
/* Specifically for the last row in tbody if we want to close it, but we usually have a footer row or spacer */

/* Remove top border for the first row of cells as it shares with header */
.custom-table th:first-child,
.custom-table td:first-child {
  border-left: none;
}

.custom-table th:last-child,
.custom-table td:last-child {
  border-right: none;
}

/* Tax Breakdown Table overrides */
.tax-breakdown-section .custom-table th,
.tax-breakdown-section .custom-table td {
  border: 1px solid #000;
}

.tax-breakdown-section .custom-table th {
  border-top: none;
  /* Already has section border */
}

.tax-breakdown-section .custom-table tr:last-child td {
  border-bottom: none;
}

.tax-breakdown-section .custom-table th:first-child,
.tax-breakdown-section .custom-table td:first-child {
  border-left: none;
}

.tax-breakdown-section .custom-table th:last-child,
.tax-breakdown-section .custom-table td:last-child {
  border-right: none;
}

.footer-section {
  font-size: 11px;
}

@media print {
  body * {
    visibility: hidden;
  }

  /* Reset layout constraints for parents if possible, though visibility handles most visibility */
  .v-dialog,
  .v-overlay,
  .v-application {
    box-shadow: none !important;
    background: none !important;
  }

  .invoice-container,
  .invoice-container * {
    visibility: visible !important;
  }

  .invoice-container {
    position: fixed !important;
    left: 0 !important;
    top: 0 !important;
    width: 100vw !important;
    height: 100vh !important;
    margin: 0 !important;
    padding: 0 !important;
    background: white !important;
    z-index: 999999 !important;
    overflow: visible !important;
    max-width: none !important;
  }

  /* Hide scrollbars */
  ::-webkit-scrollbar {
    display: none;
  }
}
</style>
```
