Feature: Unit Conversions
      I need to convert pounds to kilograms
      Scenario: convert 0 pounds to kilograms
        When i want to convert 0.0 pound into kilograms
        Then I don't expect error
      Scenario: convert some pounds to kilograms
        When i want to convert 40.0 pound into kilograms
        Then I don't expect error
      Scenario: convert minus pounds to kilograms
        When i want to convert -3.0 pound into kilograms
        Then I expect error

      Scenario: convert 0 feet to meters
        When i want to convert 0.0 feet into meters
        Then I don't expect error
      Scenario: convert minus feet to meters
        When i want to convert -10.0 feet into meters
        Then I expect error
      Scenario: convert some feet to meters
        When i want to convert 25.0 feet into meters
        Then I don't expect error

      Scenario: convert 0 fahrenheit to celsius
        When i want to convert 0.0 fahrenheit to celsius
        Then I don't expect error
      Scenario: convert minus fahrenheit to celsius
        When i want to convert -10.0 fahrenheit to celsius
        Then I don't expect error
      Scenario: convert some fahrenheit to celsius
        When i want to convert 25.0 fahrenheit to celsius
        Then I don't expect error

    Scenario: Tonne to Atomic Mass Conversion
        Given a tonne value of 2
        When I convert it to atomic mass
        Then the result should be 1.204428152e+27

      Scenario: Kilometer to Angstrom Conversion
        Given a kilometer value of 5
        When I convert it to angstrom
        Then the result should be 5e+13

      Scenario: Celsius to Planck Conversion
        Given a celsius value of 10
        When I convert it to planck
        Then the result should be 1.416808e+33

