# order_management

Sales Analytics Backend System
This repository contains a backend system designed for loading sales data, periodic data refreshing, and performing revenue analysis through a RESTful API. It demonstrates proficiency in database design, ETL (Extract, Transform, Load) pipelines, background processing, and building scalable APIs.

Table of Contents
    Objective

    Problem Statement

    Technologies Used

    Setup and Installation

    Running the Application

    API Endpoints

    Database Schema

    Data Loading and Refreshing

    ore Analysis - Revenue Calculations

    Design Decisions and Trade-offs

    Assumptions




Objective
    The objective of this project is to design a modular, scalable backend system that loads and refreshes sales data efficiently, while offering powerful revenue analysis capabilities through well-structured APIs.



Problem Statement
    Build a backend system that can:

    Load and normalize sales data from CSV into a database.

    Set up a periodic or on-demand data refresh mechanism.

    Provide APIs to trigger and retrieve various revenue calculations:

    Total Revenue

    Revenue by Product

    Revenue by Category

    Revenue by Region

Technologies Used
    Go (Golang): Backend development.

    MySQL: Database.

    GORM: ORM for database operations.

    Echo: HTTP web framework for building RESTful APIs.

    Logrus: Structured logging.

    Gocron: For periodic background jobs.

    CSV Reader: For parsing and validating data.
    Setup and Installation

Prerequisites:

    Go 1.16 or later installed.
    MySQL installed and running.
    Git installed.
    Clone the repository:

git clone [https://github.com/AbhinashKumarSingh/order_processing.git]
cd order-processing
Create the MySQL database and tables:

Create a database named orders_db.
Run the provided SQL script (schema.sql) .

Configure database connection:

    Update the database connection details in config/config.go.
    dsn := "root:{{password}}@unix(/opt/homebrew/var/mysql/mysql.sock)/orders_db?parseTime=true&tls=false"
    Replace the username, password and socket path with your local mysql setup.
    Install dependencies:

    go mod tidy
    Running the Application
    Start the application:

    go run main.go
    The application will start on port 8000.

CURL 
1. Total Revenue by region for a date range
curl --location --request POST 'http://localhost:8000/sales-service/order/revenue_by_region?start_date=2024-01-01&end_date=2024-12-31'
Response:{
    "revenue": [
        {
            "total_revenue": 2777.97,
            "region": "Asia"
        },
        {
            "total_revenue": 1299,
            "region": "Europe"
        },
        {
            "total_revenue": 349.99,
            "region": "North America"
        },
        {
            "total_revenue": 180,
            "region": "South America"
        }
    ]
}


2. Total Revenue by region for a date range

curl --location --request POST 'http://localhost:8000/sales-service/order/revenue_by_category?start_date=2024-01-01&end_date=2024-12-31'
Response:{
    "revenue": [
        {
            "total_revenue": 4606.96,
            "category": "Shoes"
        }
    ]
}


3. Total Revenue by prodcuts for a date range


 curl --location --request POST 'http://localhost:8000/sales-service/order/revenue_by_products?start_date=2024-01-01&end_date=2024-12-31'
Response:{
    "revenue": [
        {
            "total_revenue": 4604.776075,
            "product_name": "UltraBoost Running Shoes"
        }
    ]
}

Data Loading and Refreshing
    Sales data is loaded from a CSV file.

    A background job (optional) refreshes the database daily or on demand.

    Data refresh ensures no duplicates and supports overwriting or appending strategies.

    Refresh activities (success/failure) are logged for audit and debugging purposes.


Core Analysis - Revenue Calculations
    The system provides revenue analysis over a date range:

    Total Revenue: Sum of all sale amounts.

    Revenue by Product: Revenue grouped by each product.

    Revenue by Category: Revenue grouped by product categories.

    Revenue by Region: Revenue grouped by region.

Design Decisions and Trade-offs
    GORM: Chosen for faster development and type-safety in database operations.

    ocron: Lightweight library for background scheduling.

    In-app Refresh: For simplicity. In production, data pipelines or ETL orchestration tools could be better.

    Single Database: No sharding or replication considered for this prototype.

Assumptions
    Sales data provided in CSV is clean and properly formatted.

    Each order is linked to exactly one product.

    Background job (cron) is sufficient for periodic refresh without heavy load.

    Sale amount is always a positive decimal.

