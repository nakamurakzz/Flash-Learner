users/
  {userId}/
    username: string
    email: string
    password_hash: string
    created_at: timestamp
    last_login: timestamp
    subscriptions/
      {subscriptionId}/
        plan_type: string
        start_date: timestamp
        end_date: timestamp
    sentence_groups/
      {groupId}/
        title: string
        description: string
        is_public: boolean
        created_at: timestamp
        updated_at: timestamp
        sentences/
          {sentenceId}/
            text: string
            difficulty: number
            created_at: timestamp
    learning_progress/
      {sentenceId}/
        repetition_count: number
        ease_factor: number
        next_review_date: timestamp

published_groups/
  {groupId}/
    original_group_id: string (reference to sentence_groups/{groupId})
    user_id: string (reference to users/{userId})
    download_count: number
    average_rating: number
    published_at: timestamp
    reviews/
      {reviewId}/
        user_id: string (reference to users/{userId})
        rating: number
        comment: string
        created_at: timestamp