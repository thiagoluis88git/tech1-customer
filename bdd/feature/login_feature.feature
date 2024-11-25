Feature: Login
  In order to get the customer token
  As a customer of the fastfoot restaurant
  I need to be able to login with my CPF

  Scenario: then user try to pay the order, success should be displayed
    When I send "POST" request to "/login" with payload:
      """
      {
          "cpf": "12345678910"
      }   
      """
    Then the response code should be 200
    And the response payload should match json:
      """
        {
            "accessToken": "TOKEN"
        } 
      """