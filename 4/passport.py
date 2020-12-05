import re

def is_valid_birth_year(s):
  try:
    year = int(s)
  except:
    return False
  return 1920 <=year <= 2002

def is_valid_issue_year(s):
  try:
    year = int(s)
  except:
    return False
  return 2010 <=year <= 2020

def is_valid_exp_year(s):
  try:
    year = int(s)
  except:
    return False
  return 2020 <=year <= 2030

def is_valid_height(s):
  n, unit = s[:-2], s[-2:]
  try:
    n = int(n)
  except:
    return False
  if unit == 'cm':
    return 150 <= n <= 193
  if unit == 'in':
    return 59 <= n <= 76

def is_valid_hair_color(s):
  return re.match(r'^#[\da-f]{6}$', s)

def is_valid_eye_color(s):
  return s in {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

def is_valid_pid(s):
  return re.match(r'^\d{9}$', s)


REQUIRED_FIELDS = {
    "byr": is_valid_birth_year , # Birth Year
    "iyr": is_valid_issue_year,  # Issue Year
    "eyr": is_valid_exp_year,    # Expiration Year
    "hgt": is_valid_height,      # Height
    "hcl": is_valid_hair_color,  # Hair Color
    "ecl": is_valid_eye_color,   # Eye Color
    "pid": is_valid_pid,         # Passport ID
    # "cid", # Country ID
}

passports = [chunk.split() for chunk in open("input.txt").read().split("\n\n")]

def is_valid_passport(passport):
  data = dict(line.split(':')  for line in passport)
  for field, func in REQUIRED_FIELDS.items():
    if field not in data:
      return False
    if not func(data[field]):
      return False
  return True

print(len(
    [True for p in passports if is_valid_passport(p)]
))