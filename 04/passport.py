import re


def is_valid_height(s):
  n, unit = int(s[:-2]), s[-2:]
  if unit == 'cm':
    return 150 <= n <= 193
  if unit == 'in':
    return 59 <= n <= 76


REQUIRED_FIELDS = {
    "byr": lambda s: 1920 <=int(s) <= 2002, # Birth Year
    "iyr": lambda s: 2010 <=int(s) <= 2020, # Issue Year
    "eyr": lambda s: 2020 <=int(s) <= 2030, # Expiration Year
    "hgt": is_valid_height, # Height
    "hcl": lambda s: re.match(r'^#[\da-f]{6}$', s), # Hair Color
    "ecl": lambda s: s in {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, # Eye Color
    "pid": lambda s:re.match(r'^\d{9}$', s), # Passport ID
    # "cid", # Country ID
}

passports = [chunk.split() for chunk in open("input.txt").read().split("\n\n")]

def is_valid_passport(passport):
  data = dict(line.split(':')  for line in passport)
  for field, func in REQUIRED_FIELDS.items():
    if field not in data:
      return False
    try:
      if not func(data[field]):
        return False
    except:
      return False
  return True

print(len(
    [True for p in passports if is_valid_passport(p)]
))