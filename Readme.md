
# Farming Data Go Application

This project is a Farming Data Management System built using the Go programming language and MongoDB as the database. It provides a set of API endpoints to manage information related to countries, farmers, farms, schedules, and bill calculations based on fertilizer prices.










## Getting Started

### Prerequisites
1. Go (Golang)

2. MongoDB

## Installation

To get started with the Farming Data Management, follow these steps:

1. **Clone the Repository**: 
```
git clone https://github.com/swag2716/Farming-data.git
```


2. **Navigate to Project Directory**: 
```
cd farming-data
```
3. **Install dependencies:**
```
go mod download
```
4. **Configure MongoDB:**

Ensure MongoDB is running locally or update the connection string in the code accordingly.

5. ** Run the application:**
```
go run main.go
```
## API Reference

### Overview
Explore the API Documentation to gain comprehensive insights into the available API endpoints, their functionalities, and the expected responses.

1. **Get All Farms:**
- Endpoint: `/get_all_farms`
- Method: `GET`
- Action: `List All the Farms`

2. **Get Due Schedules:**

- Endpoint: `/get_due_schedules/:farmId`
- Method: `GET`
- Action: `Get schedules that are due today or tomorrow for a specific farm.`

3. **Get All Farmers Growing Crop:**

- Endpoint: `/get_all_farmers_growing_crop`
- Method: `GET`
- Action: `Fetch a list of unique farmers currently growing crops.` 

4. **Calculate Bill of Materials:**

- Endpoint: `/calculate_bill/:farmerId`
- Method: `GET`
- Action: ` Calculate the total bill for fertilizers used by a farmer based on provided fertilizer prices.`


## Deployed Version

The code is currently deployed and accessible at https://farming-data-2.onrender.com/.

