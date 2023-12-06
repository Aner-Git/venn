function hasActiveFilters(filters) {
  const keys = Object.keys(filters);
  return keys.some((k) => filters[k].active);
}

//Get an object of the active filter
function getFilters(filters) {
  let collection = {};
  const keys = Object.keys(filters);
  keys.forEach((filterName) => {
    const filter = filters[filterName];
    if (filter.active) {
      collection[filterName] = JSON.stringify(filter.export());
    }
  });
  return collection;
}

export { hasActiveFilters, getFilters };
