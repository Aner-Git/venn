function hasActiveFilters(filters) {
  const keys = Object.keys(filters);
  return keys.some((k) => filters[k].active);
}

export { hasActiveFilters };
