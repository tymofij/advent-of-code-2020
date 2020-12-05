REQUIRED_FIELDS = {
		"byr", # Birth Year
		"iyr", # Issue Year
		"eyr", # Expiration Year
		"hgt", # Height
		"hcl", # Hair Color
		"ecl", # Eye Color
		"pid", # Passport ID
		# "cid", # Country ID
}

passports = [chunk.split() for chunk in open("input.txt").read().split("\n\n")]

def field_names(passport_data):
    return set(line.split(':')[0] for line in passport_data)

print(len(
    [True for p in passports if not REQUIRED_FIELDS - field_names(p)]
))