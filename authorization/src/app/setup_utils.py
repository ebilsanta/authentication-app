import uuid

def build_url(url, params):
    ret = url + '?'
    for k in params:
        ret += k + '=' + params[k] + '&'
    return ret[:-1]

def is_uuid(test):
    try:
        uuid.UUID(str(test))
        return True
    except ValueError:
        return False