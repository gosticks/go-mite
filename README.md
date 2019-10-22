[![Go Report Card](https://goreportcard.com/badge/github.com/gosticks/go-mite)](https://goreportcard.com/report/github.com/gosticks/go-mite) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

# SimpleMite go API implementation

This library provides easy access to the mite time tracking api (API not fully covered yet...).

## Getting Started

```
    // create a mite api instance
    // appName should be a descriptive string for you application (e.g. "my-app/v0.1")
    miteAPI := NewMiteAPI(username, team, apiKey, appName)

    // now you can use the api
    customers, errCustomers := mite.GetCustomers(nil)
    if errCustomers != nil {
        l.Error(errCustomers)
        // handle error
    }
```

## Available Methods

since this library was created for a specific use case not all API endpoints are implemented yet. But I will try to have a fill API coverage asap. The library provides the following options:

### Time Entry (single)

Get single time entry

```
    entry, errEntry := miteAPI.GetTimeEntry(id)

```

Create Time entry

```
    // create a time entry instance
    entry := TimeEntry{...}

    // pass that instance to the miteAPI
    // mite will create the item in the provided userID.
    // if the userID inside cannot be written to (coworker to coworker) the entry will be created in the apiKey user.
    resp, err := miteAPI.CreateTimeEntry(entry)
```

Update Time entry

```
    update := &TimeEntry{
        // something you want to chance
    }

    // send updates to mite
    errUpdate := miteAPI.UpdateTimeEntry(id, update)
```

Delete Time entry

```
    errDelete := miteAPI.DeleteTimeEntry(id)
```

### Time Entries

Get all time entries for a time range
`filters`are a map[string]string and will be passed as GET parameters so they can be anything supported by mite
(check mite docs for more info https://mite.yo.lk/en/api/time-entries.html)

```
    // from and to are default golang time.Time instances
    entries, errEntries := miteAPI.GetTimeEntries(from, to, filters)
```

Get grouped time entries
same as get all time entries but the filters support groupBy and the result type is a `TimeEntryGroup`

```
    groups, errGroups := miteAPI.GetTimeEntriesGroup(from, to, filters)
```

### Customers

Get all customers

```
    customers, errCustomers := miteAPI.GetCustomers(nil)
```

Get customer by name

```
    customer, err := miteAPI.GetCustomerByName(customerName)
```

### Users

Get all users. When `archived` is true, the archived users will be returned.

```
    users, errUsers := miteAPI.GetUsers(archived)
```

Get a user by ID

```
    user, errUser := miteAPI.GetUser(userID)
```

### Account

Get Account

```
    account, errAcc := miteAPI.GetAccount()
```

Get Myself
Returns the current user

```
    myself, errMyself := miteAPI.GetMyself()
```

### Services

Get a single service by id

```
    service, errService := miteAPI.GetService(id)
```

Get All services

```
    services, errServices := miteAPI.GetServices(nil)
```

### Project

Get a single project by id

```
    project, errProject := miteAPI.GetProject(id)
```

Get All projects

```
    projects, errProjects := miteAPI.GetProject(nil)
```

Create Project

```
    createdProject, err := miteAPI.CreateProject(project)
```
