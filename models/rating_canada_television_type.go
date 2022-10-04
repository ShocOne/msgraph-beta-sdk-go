package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type RatingCanadaTelevisionType int

const (
    // Default value, allow all TV shows content
    ALLALLOWED_RATINGCANADATELEVISIONTYPE RatingCanadaTelevisionType = iota
    // Do not allow any TV shows content
    ALLBLOCKED_RATINGCANADATELEVISIONTYPE
    // The C classification is suitable for children ages of 2 to 7 years
    CHILDREN_RATINGCANADATELEVISIONTYPE
    // The C8 classification is suitable for children ages 8+
    CHILDRENABOVE8_RATINGCANADATELEVISIONTYPE
    // The G classification is suitable for general audience
    GENERAL_RATINGCANADATELEVISIONTYPE
    // PG, Parental Guidance
    PARENTALGUIDANCE_RATINGCANADATELEVISIONTYPE
    // The 14+ classification is intended for viewers ages 14 and older
    AGESABOVE14_RATINGCANADATELEVISIONTYPE
    // The 18+ classification is intended for viewers ages 18 and older
    AGESABOVE18_RATINGCANADATELEVISIONTYPE
)

func (i RatingCanadaTelevisionType) String() string {
    return []string{"allAllowed", "allBlocked", "children", "childrenAbove8", "general", "parentalGuidance", "agesAbove14", "agesAbove18"}[i]
}
func ParseRatingCanadaTelevisionType(v string) (interface{}, error) {
    result := ALLALLOWED_RATINGCANADATELEVISIONTYPE
    switch v {
        case "allAllowed":
            result = ALLALLOWED_RATINGCANADATELEVISIONTYPE
        case "allBlocked":
            result = ALLBLOCKED_RATINGCANADATELEVISIONTYPE
        case "children":
            result = CHILDREN_RATINGCANADATELEVISIONTYPE
        case "childrenAbove8":
            result = CHILDRENABOVE8_RATINGCANADATELEVISIONTYPE
        case "general":
            result = GENERAL_RATINGCANADATELEVISIONTYPE
        case "parentalGuidance":
            result = PARENTALGUIDANCE_RATINGCANADATELEVISIONTYPE
        case "agesAbove14":
            result = AGESABOVE14_RATINGCANADATELEVISIONTYPE
        case "agesAbove18":
            result = AGESABOVE18_RATINGCANADATELEVISIONTYPE
        default:
            return 0, errors.New("Unknown RatingCanadaTelevisionType value: " + v)
    }
    return &result, nil
}
func SerializeRatingCanadaTelevisionType(values []RatingCanadaTelevisionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
