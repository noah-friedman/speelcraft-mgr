@wip
Feature: web server
  In order to provide mods directly from the server to clients
  As a Minecraft server admin
  I need to run a web server for clients to download from

  Rule: start successfully
    Background:
      Given I create a "mods" directory
      And I create a server instance
      And I set the server path to "test"

    Scenario: start server
      When I start the server
      Then no errors should have occurred

    Scenario: start server with HTTPS
      Given I enable HTTPS for the server
      And I create a key
      And I create a certificate
      When I start the server
      Then no errors should have occurred

  Rule: fail to start server

    Scenario: start with no mods folder
      When I start the server
      Then a "speelcraft-mgr" error should have occurred

    Scenario Template: start HTTPS without key or cert
      Given I create a <file>
      When I start the server
      Then a "tls" error should have occurred

      Scenarios:
      | file        |
      | key         |
      | certificate |

