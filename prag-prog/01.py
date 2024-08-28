#!/usr/bin/python3

class Account(object):
    def __init__(self):
        self.balance = 0
        self.debits = 0
        self.credits = 0
        self.fees = 0


def format_value(value):
    if value < 0:
        return "R${:10.2f}".format(-value)
    else:
        return "R${:10.2f}".format(value)


def format_key(key, num_separators):
    return f"{key.capitalize()} {'\t' * num_separators}"


def print_separator():
    print(format_key("----------------", 1))


def print_report(key, value, num_separator):
    print(f"{format_key(key, num_separator)}{format_value(value)}")


def print_balance(account):
    print_report("debits: ", account.debits, 1)
    print_report("credits: ", account.credits, 1)
    print_report("fees: ", account.fees, 2)
    print_separator()
    print_report("balance: ", account.balance, 1)


if __name__ == '__main__':
    account = Account()
    account.debits = 100
    account.credits = 1000
    account.fees = 50
    account.balance = account.credits - account.debits - account.fees

    print_balance(account)

