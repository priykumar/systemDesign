Goal
Design an Uber-like system that supports both Rider Flow and Driver Flow with realistic behaviors such as OTP verification, surge pricing, and cab assignment

Core Functionalities
    - Rider Flow
        -   Rider registers and sets current location
        -   Rider selects destination
        -   System displays available vehicle types with estimated fares (baseddistance + vehicle type)
        -   Rider chooses vehicle type and books
        -   OTP is generated and shown to the rider
        -   Rider shares OTP with driver to start the ride
    -   Driver Flow
        -   Drivers register and link to a cab
        -   When a rider books, all nearby available drivers of selected type are notified
        -   First driver to accept gets the booking
        -   Driver enters OTP to start ride
        -   Only driver can end the ride


************ MODELS ************
Location
    - lat: float64
    - long: float64

Rider Details
    - id: int
    - name: string
    - phoneNumber: string
    - location: Location

Cab Details
    - id: int
    - cabNumber: string
    - vahicleType: const
    - isAvailable: bool

Booking Details:
    - id: int
    - rider: Rider
    - cab: Cab
    - otp: int
    - pickUpLocation: Location
    - dropLocation: Location
    - fare: float64
