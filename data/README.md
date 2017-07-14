# Data

The data is structured in the follow way:

   - First line represents how many routes there are.
   - First number of each line represents the `routeId`
   - Each number separed by space is the `stopId`

Example:

```
2
1 3 5 6 2
2 5 10 4
```

The data above is the same as this `json`:
```
{
  'routes': [
    { 'id': 1, stops: [ 3, 5, 6, 2 ] },
    { 'id': 2, stops: [ 5, 10, 4 ] }
  ]
}
```
