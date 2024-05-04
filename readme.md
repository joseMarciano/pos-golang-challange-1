### 🌤️ Weather Service README 🌡️

---

#### Objective:
To develop a system in Go that receives a ZIP code, identifies the city, and returns the current weather (temperature in Celsius, Fahrenheit, and Kelvin).

#### Requirements:
- The system must receive a valid 8-digit ZIP code.
- It should search for the ZIP code and find the location name, then return the temperatures formatted as Celsius, Fahrenheit, and Kelvin.
- The system should respond appropriately in the following scenarios:
    - **Success Case:**
        - HTTP Code: 200
        - Response Body: `{ "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }`
    - **Failure Case, for invalid ZIP code format:**
        - HTTP Code: 422
        - Message: invalid zipcode
    - **Failure Case, if the ZIP code is not found:**
        - HTTP Code: 404
        - Message: cannot find zipcode

#### Tips:
- Use the ViaCEP API (or similar) to find the location for temperature querying: [ViaCEP API](https://viacep.com.br/)
- Use the WeatherAPI (or similar) to fetch the desired temperatures: [WeatherAPI](https://www.weatherapi.com/)
- To convert Celsius to Fahrenheit, use the formula: F = C * 1.8 + 32
- To convert Celsius to Kelvin, use the formula: K = C + 273
    - Where F = Fahrenheit, C = Celsius, and K = Kelvin.

#### Delivery:
- Complete source code of the implementation.
- Automated tests demonstrating functionality.
- Utilize docker/docker-compose for testing purposes.
