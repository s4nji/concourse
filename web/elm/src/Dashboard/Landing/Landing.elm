module Dashboard.Landing.Landing exposing
    ( Model
    , view
    )

import Concourse
import HoverState
import Html exposing (Html)
import Html.Attributes exposing (href, id, class, title)
import Html.Events exposing (onClick, onMouseEnter, onMouseLeave)
import List.Extra
import Message.Callback exposing (Callback(..))
import Message.Message exposing (DomID(..), Message(..))
import Message.Subscription exposing (Delivery(..))
import RemoteData exposing (RemoteData(..), WebData)
import Set exposing (Set)
import Dashboard.Landing.Styles as Styles
import Views.Icon as Icon
import Routes

type alias Model m =
    { m
        | expandedTeams : Set String
        , pipelines : WebData (List Concourse.Pipeline)
        , hovered : HoverState.HoverState
    }

view : Model m -> Html Message
view model =
    Html.div
        [class "landing"]
        (model.pipelines
            |> RemoteData.withDefault []
            |> List.sortBy .teamName
            |> List.Extra.gatherEqualsBy .teamName
            |> List.map
                (\(p, ps) ->
                    team
                        { hovered = model.hovered
                        , isExpanded = Set.member p.teamName model.expandedTeams
                        , teamName = p.teamName
                        , pipelines = p :: ps
                        }
                )
        )

team :
    { a
        | hovered : HoverState.HoverState
        , isExpanded : Bool
        , teamName : String
        , pipelines : List (Concourse.Pipeline)
    }
    -> Html Message
team ({ isExpanded, pipelines } as session) =
    Html.div
        Styles.team
        [ teamHeader session
        , if isExpanded then
            Html.div Styles.column
              ( pipelines
                |> List.sortBy .name
                |> List.map (renderPipeline session)
              )
          else
            Html.text ""
        ]

teamHeader :
    { a
        | hovered : HoverState.HoverState
        , isExpanded : Bool
        , teamName : String
    }
    -> Html Message
teamHeader { hovered, isExpanded, teamName } =
    let
        isHovered =
            HoverState.isHovered
                (SideBarTeam teamName)
                hovered

    in
    Html.div
        (Styles.teamHeader
            ++ [ onClick <| Click <| SideBarTeam teamName
               , onMouseEnter <| Hover <| Just <| SideBarTeam teamName
               , onMouseLeave <| Hover Nothing
               ]
        )
        [ Styles.teamIcon { isHovered = isHovered }
        , Styles.arrow
            { isHovered = isHovered
            , isExpanded = isExpanded
            }
        , Html.div
            (title teamName
                :: Styles.teamName { isHovered = isHovered }
            )
            [ Html.text teamName ]
        ]

renderPipeline :
    { a
        | hovered : HoverState.HoverState
        , teamName : String
    }
    -> Concourse.Pipeline
    -> Html Message
renderPipeline { hovered, teamName } p =
    let
        pipelineId =
            { pipelineName = p.name
            , teamName = teamName
            }

        isHovered =
            HoverState.isHovered
                (SideBarPipeline pipelineId)
                hovered
    in
    Html.div Styles.pipeline
        [ Html.div
            (Styles.pipelineIcon { isHovered = isHovered }
            )
            []
        , Html.a
            (Styles.pipelineLink { isHovered = isHovered }
                ++ [ href <|
                        Routes.toString <|
                            Routes.Pipeline { id = pipelineId, groups = [] }
                   , title p.name
                   , onMouseEnter <| Hover <| Just <| SideBarPipeline pipelineId
                   , onMouseLeave <| Hover Nothing
                   ]
            )
            [ Html.text p.name ]
        ]
