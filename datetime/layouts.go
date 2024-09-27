package datetime

// Common layouts for formatting and parsing date and time strings.
const (
	fullTimeWithMonotonicClockLayout 	= "2006-01-02 15:04:05.999999999 -0700 MST m=+0.000000001"
	fullTimeLayoutTZ      				= "2006-01-02 15:04:05.999999999 -0700 MST"  // Full timestamp with nanoseconds and time zone.
	dateTimeLayoutTZ     			 	= "2006-01-02 15:04:05 -0700 MST"           // Date and time with time zone.
	dateTimeLayout        				= "2006-01-02 15:04:05"                     // Date and time without time zone.
	dateLayout            				= "2006-01-02"                              // Date only (YYYY-MM-DD).
	yearLayout            				= "2006"                                    // Year (YYYY).
	dayLayout            			 	= "02"                                      // Day of the month (DD).
	hourLayout            				= "15:04:05"                                // Hour (HH:MM:SS).
	secondsLayout         				= "05"                                      // Seconds (SS).
	millisecondsLayout    				= "15:04:05.000"                            // Time with milliseconds.
	fullDateTimeWithTZ    				= "2006-01-02 15:04:05.999999999....-0700....MST"  // Full date-time with custom separators.
)
