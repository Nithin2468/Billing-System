<template>
  <v-container fluid class="pa-6">
    <v-card elevation="2" rounded="lg">
      <v-card-title
        class="bg-teal-darken-1 text-white py-4 px-6 d-flex align-center"
      >
        <v-icon icon="mdi-chart-box" class="mr-3"></v-icon>
        <span class="text-h5 font-weight-bold">Stock & Sales Report</span>
        <v-spacer></v-spacer>
        <v-btn
          color="white"
          variant="elevated"
          prepend-icon="mdi-microsoft-excel"
          class="text-teal-darken-1 font-weight-bold"
          @click="downloadExcel"
          :disabled="!reports.length"
        >
          Download Excel
        </v-btn>
      </v-card-title>

      <v-card-text class="pa-6">
        <!-- Filters -->
        <v-row align="center" class="mb-4">
          <v-col cols="12" md="4">
            <v-text-field
              v-model="search"
              label="Search Medicine"
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              density="comfortable"
              hide-details
            ></v-text-field>
          </v-col>
          <v-col cols="12" md="3">
            <v-text-field
              v-model="startDate"
              label="From Date"
              type="date"
              variant="outlined"
              density="comfortable"
              hide-details
            ></v-text-field>
          </v-col>
          <v-col cols="12" md="3">
            <v-text-field
              v-model="endDate"
              label="To Date"
              type="date"
              variant="outlined"
              density="comfortable"
              hide-details
            ></v-text-field>
          </v-col>
          <v-col cols="12" md="2">
            <v-btn
              color="teal"
              variant="elevated"
              block
              height="48"
              @click="fetchReport"
              :loading="loading"
            >
              Apply Filter
            </v-btn>
          </v-col>
        </v-row>

        <!-- Data Table -->
        <v-data-table
          :headers="headers"
          :items="reports"
          :search="search"
          class="elevation-0 border"
          hover
          :loading="loading"
        >
          <template v-slot:[`item.currentQty`]="{ item }">
            <v-chip
              :color="item.currentQty < 10 ? 'error' : 'success'"
              size="small"
              class="font-weight-bold"
              label
            >
              {{ item.currentQty }}
            </v-chip>
          </template>
          <template v-slot:[`item.totalSales`]="{ item }">
            <span class="font-weight-bold text-teal-darken-2">
              ₹{{ item.totalSales.toFixed(2) }}
            </span>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import EventService from "../services/eventServices";

export default {
  name: "ManagerStockReport",
  data() {
    return {
      search: "",
      startDate: "",
      endDate: "",
      loading: false,
      reports: [],
      headers: [
        { title: "Medicine Name", key: "name", align: "start" },
        { title: "Batch ID", key: "batchId", align: "start" },
        { title: "Expiry Date", key: "expiryDate", align: "start" },
        { title: "Current Stock", key: "currentQty", align: "center" },
        {
          title: "Sold Qty (Selected Period)",
          key: "soldQty",
          align: "center",
        },
        { title: "Total Sales (₹)", key: "totalSales", align: "end" },
      ],
    };
  },
  mounted() {
    // Set default date range (Current Month)
    const now = new Date();
    const firstDay = new Date(now.getFullYear(), now.getMonth(), 1);
    this.startDate = firstDay.toISOString().substr(0, 10);
    this.endDate = now.toISOString().substr(0, 10);

    this.fetchReport();
  },
  methods: {
    async fetchReport() {
      this.loading = true;
      try {
        const params = {
          startDate: this.startDate,
          endDate: this.endDate,
        };
        const response = await EventService.getStockReport(params);
        this.reports = response.data || [];
      } catch (error) {
        console.error("Error fetching stock report:", error);
        // alert("Failed to fetch report");
      } finally {
        this.loading = false;
      }
    },
    downloadExcel() {
      if (!this.reports.length) return;

      // Create CSV Content
      const headers = [
        "Medicine Name,Batch ID,Expiry Date,Current Stock,Sold Qty,Total Sales",
      ];
      const rows = this.reports.map(
        (item) =>
          `"${item.name}","${item.batchId}","${item.expiryDate}",${
            item.currentQty
          },${item.soldQty},${item.totalSales.toFixed(2)}`
      );

      const csvContent =
        "data:text/csv;charset=utf-8," + [headers, ...rows].join("\n");
      const encodedUri = encodeURI(csvContent);

      // Create Download Link
      const link = document.createElement("a");
      link.setAttribute("href", encodedUri);
      link.setAttribute(
        "download",
        `Stock_Sales_Report_${this.startDate}_to_${this.endDate}.csv`
      );
      document.body.appendChild(link);

      link.click();
      document.body.removeChild(link);
    },
  },
};
</script>
