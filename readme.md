# Space Tickets

## Contents
<!-- `make toc` to generate https://github.com/jonschlinkert/markdown-toc#cli -->

<!-- toc -->

- [To Run](#to-run)
- [Valid Schedules](#valid-schedules)
  * [Example Requests](#example-requests)
- [Improvements](#improvements)

<!-- tocstop -->

## To Run
You need Docker on your machine for this to work.

It is assumed you're on a Mac.

Run `make start logs` from the project folder in a terminal.

This will pull the Docker image for the application, run it, and show its logs.

If the image won't pull down from Docker Hub, you can build it locally by running `make build`. This builds a multi-platform image that will work on `amd64` and `arm64` machines. It uses Docker's `buildx` tool, which should be installed on your machine if you have Docker Desktop.

Following the build, run the app using `make start logs`.

Access the API here http://localhost:8080/api/v1.

View Swagger docs here http://localhost:8081.

To run unit tests run `make test`.

## Valid Schedules

SpaceX data from https://api.spacexdata.com ends on 1st December 2022, so anything after then will not clash with a SpaceX launch.

Here's the launchpad schedule so you know what flights are valid. You will still need to know what day of the week your desired launch date falls on. This [site](https://www.calculator.net/day-of-the-week-calculator.html) can help with that.

| Launchpad                                                    | launchpad_id                             | Destination   | destination_id                             | Day of Week |
|--------------------------------------------------------------|------------------------------------------|---------------|--------------------------------------------|-------------|
| Cape Canaveral  Launch Complex 40   | b542c0cf-7fe3-4bb1-a63f-7cbdf8359975     | Moon          | 466fc378-14eb-4ed9-8bec-d29abe54c5a9      | Monday      |
|   |     | Mars          | f47eef79-675f-46da-86f9-ee598185d204      | Tuesday     |
|   |     | Pluto         | fbd40165-03c7-47a5-be72-c79f81ebbf67      | Wednesday   |
|    |     | Asteroid Belt | 13b91e0c-cdb4-4108-9c48-5a49d8ded732      | Thursday    |
|    |     | Europa        | 998f4a82-5a1c-4542-8497-e3fa24618d79      | Friday      |
|    |     | Titan         | 12549fca-d086-4e9f-b14e-dcb3b0d09c63      | Saturday    |
|   |     | Ganymede      | 3840d5ce-b939-4af7-9dd8-ac12c09d1493      | Sunday      |
| Kennedy Space Center Launch Complex 39A             | 4079f070-3e58-4e61-8af7-05c8de8e1fbf     | Asteroid Belt | 13b91e0c-cdb4-4108-9c48-5a49d8ded732      | Monday      |
|             |     | Europa        | 998f4a82-5a1c-4542-8497-e3fa24618d79      | Tuesday     |
|              |     | Titan         | 12549fca-d086-4e9f-b14e-dcb3b0d09c63      | Wednesday   |
|             |    | Ganymede      | 3840d5ce-b939-4af7-9dd8-ac12c09d1493      | Thursday    |
|             |     | Moon          | 466fc378-14eb-4ed9-8bec-d29abe54c5a9      | Friday      |
|            |     | Mars          | f47eef79-675f-46da-86f9-ee598185d204      | Saturday    |
|             |     | Pluto         | fbd40165-03c7-47a5-be72-c79f81ebbf67      | Sunday      |
| Kwajalein Atoll Omelek Island                                | e169113a-ae89-4c39-9a28-3cbc1c96e5e0     | Titan         | 12549fca-d086-4e9f-b14e-dcb3b0d09c63      | Monday      |
|                              |    | Ganymede      | 3840d5ce-b939-4af7-9dd8-ac12c09d1493      | Tuesday     |
|                              |     | Moon          | 466fc378-14eb-4ed9-8bec-d29abe54c5a9      | Wednesday   |
|                               |     | Mars          | f47eef79-675f-46da-86f9-ee598185d204      | Thursday    |
|                           |      | Pluto         | fbd40165-03c7-47a5-be72-c79f81ebbf67      | Friday      |
|                               |     | Asteroid Belt | 13b91e0c-cdb4-4108-9c48-5a49d8ded732      | Saturday    |
|                              |      | Europa        | 998f4a82-5a1c-4542-8497-e3fa24618d79      | Sunday      |
| SpaceX South Texas Launch Site                               | b09e0b80-51ca-44ac-820a-d5b95b209cad     | Ganymede      | 3840d5ce-b939-4af7-9dd8-ac12c09d1493      | Monday      |
|                           |     | Moon          | 466fc378-14eb-4ed9-8bec-d29abe54c5a9      | Tuesday     |
|                              |     | Mars          | f47eef79-675f-46da-86f9-ee598185d204      | Wednesday   |
|                               |   | Pluto         | fbd40165-03c7-47a5-be72-c79f81ebbf67      | Thursday    |
|                               |    | Asteroid Belt | 13b91e0c-cdb4-4108-9c48-5a49d8ded732      | Friday      |
|                              |   | Europa        | 998f4a82-5a1c-4542-8497-e3fa24618d79      | Saturday    |
|                               |     | Titan         | 12549fca-d086-4e9f-b14e-dcb3b0d09c63      | Sunday      |
| Vandenberg Space Launch Complex 3W          | d95c83bb-be3f-4bdb-93fe-77015d95f759     | Mars          | f47eef79-675f-46da-86f9-ee598185d204      | Monday      |
|          |     | Pluto         | fbd40165-03c7-47a5-be72-c79f81ebbf67      | Tuesday     |
|          |    | Asteroid Belt | 13b91e0c-cdb4-4108-9c48-5a49d8ded732      | Wednesday   |
|         |    | Europa        | 998f4a82-5a1c-4542-8497-e3fa24618d79      | Thursday    |
|           |     | Titan         | 12549fca-d086-4e9f-b14e-dcb3b0d09c63      | Friday      |
|        |    | Ganymede      | 3840d5ce-b939-4af7-9dd8-ac12c09d1493      | Saturday    |
|          |    | Moon          | 466fc378-14eb-4ed9-8bec-d29abe54c5a9      | Sunday      |
| Vandenberg Space Launch Complex 4E      | 9f8cb517-ca3b-4810-baef-80b48b8cf5e6   | Europa        | 998f4a82-5a1c-4542-8497-e3fa24618d79  | Monday      |
|                                                          |                                        | Titan         | 12549fca-d086-4e9f-b14e-dcb3b0d09c63  | Tuesday     |
|                                                          |                                        | Ganymede      | 3840d5ce-b939-4af7-9dd8-ac12c09d1493  | Wednesday   |
|                                                          |                                        | Moon          | 466fc378-14eb-4ed9-8bec-d29abe54c5a9  | Thursday    |
|                                                          |                                        | Mars          | f47eef79-675f-46da-86f9-ee598185d204  | Friday      |
|                                                          |                                        | Pluto         | fbd40165-03c7-47a5-be72-c79f81ebbf67  | Saturday    |
|                                                          |                                        | Asteroid Belt | 13b91e0c-cdb4-4108-9c48-5a49d8ded732  | Sunday      |

### Example Requests

This will successfully create a booking.

```
curl --location 'localhost:8080/api/v1/booking' \
--header 'Content-Type: application/json' \
--data '{
  "first_name": "Ian",
  "last_name": "Thomson",
  "gender": "Male",
  "birthday": "2000-04-12",
  "launch_pad_id": "b542c0cf-7fe3-4bb1-a63f-7cbdf8359975",
  "destination_id": "466fc378-14eb-4ed9-8bec-d29abe54c5a9",
  "launch_date": "2010-12-06"
}'
```

This will fail because the launchpad doesn't fly to the destination on the selected day.

```
curl --location 'localhost:8080/api/v1/booking' \
--header 'Content-Type: application/json' \
--data '{
  "first_name": "Ian",
  "last_name": "Thomson",
  "gender": "Male",
  "birthday": "2000-04-12",
  "launch_pad_id": "b542c0cf-7fe3-4bb1-a63f-7cbdf8359975",
  "destination_id": "466fc378-14eb-4ed9-8bec-d29abe54c5a9",
  "launch_date": "2010-12-07"
}'
```

This will fail because there is a SpaceX launch from that launchpad on the selected day.

```
curl --location 'localhost:8080/api/v1/booking' \
--header 'Content-Type: application/json' \
--data '{
  "first_name": "Ian",
  "last_name": "Thomson",
  "gender": "Male",
  "birthday": "2000-04-12",
  "launch_pad_id": "4079f070-3e58-4e61-8af7-05c8de8e1fbf",
  "destination_id": "fbd40165-03c7-47a5-be72-c79f81ebbf67",
  "launch_date": "2022-10-05"
}'
```

## Improvements
* Improved error messages including the launchpad name, destination name and day of the week for their desired launch data. This would help users verify what they sent
* Prevent creation of duplicate flights
* When deleting a booking, respond with error if already marked as deleted
* Cache flight schedule rather than getting it for every request
* Return user-friendly launchpad and destination names along with the launchpad and destination ids in a booking
* Validate user input like sensible birthday, date formats, recognised launch_pad_id and destination_id
* Endpoints to list the launchpads and destinations ids and names
* Store users in database so they don't have to provide their name and birthday, they could just send an id, or log in so the system knows who they are