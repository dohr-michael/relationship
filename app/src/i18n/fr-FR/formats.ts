export default {
    date    : {
        'short': {
            'day'  : 'numeric',
            'month': 'numeric',
            'year' : 'numeric'
        },
        long: {
            day     : "numeric",
            month   : "long",
            year    : "numeric"
        },
        dayMonth: {
            day     : "numeric",
            month   : "long"
        },
        monthYear: {
            month   : "numeric",
            year     : "numeric"
        },
        year: {
            year    : "numeric"
        },
        month: {
            month   : "long"
        }
    },
    time    : {
        'hhmm': {
            'hour'  : 'numeric',
            'minute': 'numeric'
        }
    },
    number  : {
        USD: {
            style                : 'currency',
            currency             : 'USD',
            minimumFractionDigits: 2
        },
        
        InvoiceNumber: {
            style                   : 'decimal',
            minimumFractionDigits   : 2
        }
    },
    relative: {
        'hours': {
            'units': 'hour',
            'style': 'numeric'
        }
    }
}
